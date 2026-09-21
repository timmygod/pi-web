package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	"pi-web/internal/chat"
	"pi-web/internal/schedules"
	"pi-web/internal/sessions"
)

// scheduleTickInterval is how often the scheduler re-evaluates due schedules.
// Cron is minute-resolution, so a sub-minute tick keeps firing within a few
// seconds of the target time.
const scheduleTickInterval = 30 * time.Second

// scheduleWorkerTimeout bounds the EnsureWorker step when a schedule fires.
const scheduleWorkerTimeout = 60 * time.Second

// scheduleState tracks, per schedule, the next time it should fire and the
// cron/timezone signature it was computed from (so edits force a recompute).
type scheduleState struct {
	next time.Time
	sig  string
}

func scheduleSig(sc schedules.Schedule) string {
	return sc.CronExpr + "|" + sc.RunAt + "|" + sc.Timezone
}

// runScheduler ticks until stopped, firing any schedule whose next occurrence
// has arrived. Missed occurrences (while the process was down) are skipped:
// a schedule's first evaluation computes its next fire from now, never the past.
func (s *Server) runScheduler(stop <-chan struct{}, interval time.Duration) {
	state := make(map[string]scheduleState)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		s.evaluateSchedules(state)
		select {
		case <-ticker.C:
		case <-stop:
			return
		}
	}
}

func (s *Server) evaluateSchedules(state map[string]scheduleState) {
	if s.schedules == nil {
		return
	}
	list, err := s.schedules.List()
	if err != nil {
		fmt.Fprintf(os.Stderr, "scheduler: list schedules: %v\n", err)
		return
	}
	now := s.now()
	seen := make(map[string]bool, len(list))
	for _, sc := range list {
		seen[sc.ID] = true
		if !sc.Enabled || (sc.IsManual() && !sc.IsOneShot()) {
			delete(state, sc.ID)
			continue
		}
		sig := scheduleSig(sc)
		st, ok := state[sc.ID]
		if !ok || st.sig != sig {
			var next time.Time
			var err error
			if sc.IsOneShot() {
				next, err = time.Parse(time.RFC3339, sc.RunAt)
			} else {
				next, err = schedules.NextFire(sc.CronExpr, sc.Timezone, now)
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "scheduler: %q invalid schedule time: %v\n", sc.Name, err)
				delete(state, sc.ID)
				continue
			}
			state[sc.ID] = scheduleState{next: next, sig: sig}
			st = scheduleState{next: next, sig: sig}
			if !sc.IsOneShot() {
				continue
			}
		}
		if now.Before(st.next) {
			continue
		}
		if sc.IsOneShot() {
			claimed, err := s.schedules.DisableOneShot(sc.ID)
			if err != nil || !claimed {
				if err != nil {
					fmt.Fprintf(os.Stderr, "scheduler: claim one-shot %q: %v\n", sc.Name, err)
				}
				delete(state, sc.ID)
				continue
			}
			sc.Enabled = false
			s.notifySchedulesChanged("disabled", sc)
		}
		sc := sc
		s.startTask(func(ctx context.Context) {
			if _, err := s.fireScheduleContext(ctx, sc); err != nil && !errors.Is(err, context.Canceled) {
				fmt.Fprintf(os.Stderr, "scheduler: fire %q: %v\n", sc.Name, err)
			}
		})
		if sc.IsOneShot() {
			delete(state, sc.ID)
			continue
		}
		next, err := schedules.NextFire(sc.CronExpr, sc.Timezone, now)
		if err != nil {
			delete(state, sc.ID)
			continue
		}
		state[sc.ID] = scheduleState{next: next, sig: sig}
	}
	for id := range state {
		if !seen[id] {
			delete(state, id)
		}
	}
}

// scheduleNameForSession reports whether a session was created by a schedule,
// returning the schedule's name. Used to route schedule-specific notifications.
func (s *Server) scheduleNameForSession(sessionID string) (string, bool) {
	if s.schedules == nil {
		return "", false
	}
	return s.schedules.ScheduleNameForSession(sessionID)
}

// fireSchedule creates a fresh pi session for the schedule, records the run, and
// sends the instructions as the first message so pi runs autonomously. Returns
// the created session's UUID. Used by both the timer and the Run-now endpoint.
func (s *Server) fireSchedule(sc schedules.Schedule) (string, error) {
	return s.fireScheduleContext(context.Background(), sc)
}

func (s *Server) fireScheduleContext(ctx context.Context, sc schedules.Schedule) (string, error) {
	if s.schedules == nil {
		return "", errors.New("schedules unavailable")
	}
	fired := s.now().UTC()
	runID, err := s.schedules.RecordRun(schedules.Run{
		ScheduleID: sc.ID,
		FiredAt:    fired.Format(time.RFC3339),
		Status:     schedules.RunStatusRunning,
	})
	if err != nil {
		return "", fmt.Errorf("record run: %w", err)
	}
	_ = s.schedules.SetLastRun(sc.ID, fired)

	path := strings.TrimSpace(sc.ProjectPath)
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			s.failScheduleRun(sc, runID, err.Error())
			return "", err
		}
		path = home
	}

	modelCandidates := s.scheduleModelCandidates(ctx, sc)
	settings := sessions.InitialSettings{
		ModelProvider: sc.ModelProvider,
		ModelID:       sc.ModelID,
		ThinkingLevel: sc.ThinkingLevel,
	}
	if len(modelCandidates) > 0 {
		settings.ModelProvider = modelCandidates[0].Provider
		settings.ModelID = modelCandidates[0].ID
	}
	filename, err := sessions.CreateSessionFileWithSettings(s.sessionsDir, path, settings)
	if err != nil {
		s.failScheduleRun(sc, runID, err.Error())
		return "", fmt.Errorf("create session: %w", err)
	}
	resolved, err := sessions.ResolveByID(s.sessionsDir, filename)
	if err != nil {
		s.failScheduleRun(sc, runID, err.Error())
		return "", fmt.Errorf("resolve session: %w", err)
	}
	sessionID := resolved.Session.ID
	if err := s.schedules.AttachSession(runID, sessionID, filename); err != nil {
		fmt.Fprintf(os.Stderr, "scheduler: attach session: %v\n", err)
	}

	if s.chatSender == nil {
		s.failScheduleRun(sc, runID, "chat unavailable")
		return sessionID, errors.New("chat unavailable")
	}
	mode, err := s.saveSessionMode(sessionID, sessionModeAuto,
		s.resolveModelMetadata(ctx, settings.ModelProvider, settings.ModelID))
	if err != nil {
		s.failScheduleRun(sc, runID, err.Error())
		return sessionID, fmt.Errorf("save session mode: %w", err)
	}
	if err := s.prepareWorkerMode(sessionID, mode); err != nil {
		s.failScheduleRun(sc, runID, err.Error())
		return sessionID, fmt.Errorf("prepare worker mode: %w", err)
	}
	workerCtx, cancel := context.WithTimeout(ctx, scheduleWorkerTimeout)
	defer cancel()
	if err := s.chatSender.EnsureWorker(workerCtx, sessionID, resolved.Path); err != nil {
		s.failScheduleRun(sc, runID, err.Error())
		return sessionID, fmt.Errorf("ensure worker: %w", err)
	}
	// Empty sessions skip history restore, so implicit file entries are ignored.
	if len(modelCandidates) > 0 {
		var setErr error
		for _, candidate := range modelCandidates {
			setErr = s.chatSender.SetModel(ctx, sessionID, resolved.Path, candidate.Provider, candidate.ID)
			if setErr == nil {
				settings.ModelProvider, settings.ModelID = candidate.Provider, candidate.ID
				_ = s.schedules.SetRunModel(runID, candidate.Provider, candidate.ID)
				_, _ = s.saveSessionMode(sessionID, sessionModeAuto, s.resolveModelMetadata(ctx, candidate.Provider, candidate.ID))
				break
			}
		}
		if setErr != nil {
			s.failScheduleRun(sc, runID, setErr.Error())
			return sessionID, fmt.Errorf("set model: no candidate worked: %w", setErr)
		}
	}
	if sc.ThinkingLevel != "" {
		if err := s.chatSender.SetThinkingLevel(ctx, sessionID, resolved.Path, sc.ThinkingLevel); err != nil {
			s.failScheduleRun(sc, runID, err.Error())
			return sessionID, fmt.Errorf("set thinking level: %w", err)
		}
	}
	if err := s.sendSessionChat(ctx, resolved, chat.Request{Message: sc.Instructions}); err != nil {
		s.failScheduleRun(sc, runID, err.Error())
		return sessionID, fmt.Errorf("send: %w", err)
	}
	return sessionID, nil
}

func (s *Server) failScheduleRun(sc schedules.Schedule, runID int64, message string) {
	_ = s.schedules.FailRun(runID, message)
	s.deliverScheduleCallback(context.Background(), sc, runID)
}

type scheduleModel struct {
	Provider, ID, Name, BaseURL string
	Starred                     bool
}

func (s *Server) scheduleModelCandidates(ctx context.Context, sc schedules.Schedule) []scheduleModel {
	selector := strings.TrimSpace(sc.ModelSelector)
	if selector == "" {
		if sc.ModelProvider == "" || sc.ModelID == "" {
			return nil
		}
		return []scheduleModel{{Provider: sc.ModelProvider, ID: sc.ModelID}}
	}
	if s.models == nil {
		return nil
	}
	raw, err := s.models(ctx)
	if err != nil {
		return nil
	}
	var payload struct {
		Models []struct {
			Provider string `json:"provider"`
			ID       string `json:"id"`
			ModelID  string `json:"modelId"`
			Name     string `json:"name"`
			BaseURL  string `json:"baseUrl"`
			Starred  bool   `json:"starred"`
		} `json:"models"`
	}
	if json.Unmarshal(raw, &payload) != nil {
		return nil
	}
	models := make([]scheduleModel, 0, len(payload.Models))
	for _, item := range payload.Models {
		id := item.ID
		if id == "" {
			id = item.ModelID
		}
		if id == "" || item.Provider == "" {
			continue
		}
		models = append(models, scheduleModel{item.Provider, id, item.Name, item.BaseURL, item.Starred})
	}
	usage, recent := map[string]int{}, map[string]int{}
	if s.cache != nil {
		if summaries, err := s.loadSummaries(); err == nil {
			sessions.SortSummariesByActivity(summaries)
			for i, summary := range summaries {
				key := summary.ModelProvider + "\x00" + summary.Model
				if summary.Model != "" {
					usage[key]++
					if _, ok := recent[key]; !ok {
						recent[key] = i
					}
				}
			}
		}
	}
	isLocal := func(m scheduleModel) bool {
		p := strings.ToLower(m.Provider)
		if p == "local" || p == "ollama" || p == "llama" || p == "lmstudio" || p == "lm-studio" {
			return true
		}
		u, err := url.Parse(m.BaseURL)
		if err != nil {
			return false
		}
		host := u.Hostname()
		ip := net.ParseIP(host)
		return host == "localhost" || (ip != nil && (ip.IsLoopback() || ip.IsPrivate()))
	}
	q := strings.ToLower(selector)
	score := func(m scheduleModel) int {
		key := m.Provider + "\x00" + m.ID
		value := usage[key] * 100
		if i, ok := recent[key]; ok {
			value += max(0, 50-i)
		}
		if m.Starred {
			value += 10000
		}
		if q == "local" && isLocal(m) {
			value += 100000
		}
		if q != "local" && q != "free" {
			id, name, provider := strings.ToLower(m.ID), strings.ToLower(m.Name), strings.ToLower(m.Provider)
			switch {
			case q == id || q == provider+"/"+id:
				value += 100000
			case q == name:
				value += 90000
			case strings.Contains(id, q) || strings.Contains(name, q):
				value += 70000
			case strings.Contains(q, id):
				value += 50000
			}
		}
		return value
	}
	if q == "local" {
		local := make([]scheduleModel, 0, len(models))
		remote := make([]scheduleModel, 0, len(models))
		for _, m := range models {
			if isLocal(m) {
				local = append(local, m)
			} else {
				remote = append(remote, m)
			}
		}
		models = append(local, remote...)
	}
	sort.SliceStable(models, func(i, j int) bool { return score(models[i]) > score(models[j]) })
	return models
}
