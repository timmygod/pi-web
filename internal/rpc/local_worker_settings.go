package rpc

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
)

type WorkerOptions struct {
	AgentDir           string
	LocalContextWindow int
}

func prepareLocalWorkerAgentDir(agentDir, sessionPath string, contextWindow int) (string, error) {
	if agentDir == "" || contextWindow <= 0 {
		return "", fmt.Errorf("local worker requires an agent directory and context window")
	}
	name := filepath.Base(sessionPath)
	name = name[:len(name)-len(filepath.Ext(name))]
	workerDir := filepath.Join(agentDir, "pi-web", "local-workers", name)
	if err := os.MkdirAll(workerDir, 0700); err != nil {
		return "", err
	}

	entries, err := os.ReadDir(agentDir)
	if err != nil {
		return "", err
	}
	for _, entry := range entries {
		switch entry.Name() {
		case "settings.json", "extensions", "sessions", "session-status", "pi-web", "pi-web.sqlite", "pi-web.sqlite-shm", "pi-web.sqlite-wal":
			continue
		}
		target := filepath.Join(workerDir, entry.Name())
		if _, err := os.Lstat(target); err == nil {
			continue
		}
		if err := os.Symlink(filepath.Join(agentDir, entry.Name()), target); err != nil && !os.IsExist(err) {
			return "", err
		}
	}
	if err := prepareLocalWorkerExtensions(agentDir, workerDir); err != nil {
		return "", err
	}

	settings := make(map[string]any)
	if data, err := os.ReadFile(filepath.Join(agentDir, "settings.json")); err == nil {
		if err := json.Unmarshal(data, &settings); err != nil {
			return "", fmt.Errorf("parse settings.json for Local Mode: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return "", err
	}
	compaction, _ := settings["compaction"].(map[string]any)
	if compaction == nil {
		compaction = make(map[string]any)
	}
	compaction["enabled"] = true
	// Pi's threshold is strictly `contextTokens > contextWindow-reserveTokens`.
	// Add one token so reaching exactly 65% satisfies that comparison.
	reserveTokens := contextWindow - int(math.Ceil(float64(contextWindow)*0.65)) + 1
	compaction["reserveTokens"] = reserveTokens
	maximumKeepRecent := contextWindow / 5
	if maximumKeepRecent < 1 {
		maximumKeepRecent = 1
	}
	keepRecent := maximumKeepRecent
	if configured, ok := numberValue(compaction["keepRecentTokens"]); ok && configured < keepRecent {
		keepRecent = configured
	}
	if keepRecent < 1 {
		keepRecent = 1
	}
	compaction["keepRecentTokens"] = keepRecent
	settings["compaction"] = compaction
	if err := validateProjectCompactionSettings(sessionPath, reserveTokens, maximumKeepRecent); err != nil {
		return "", err
	}
	encoded, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return "", err
	}
	encoded = append(encoded, '\n')
	if err := os.WriteFile(filepath.Join(workerDir, "settings.json"), encoded, 0600); err != nil {
		return "", err
	}
	return workerDir, nil
}

const localCompactionGuardFilename = "pi-web-local-compaction-guard.js"

//go:embed local_compaction_guard.mjs
var localCompactionGuardSource string

func prepareLocalWorkerExtensions(agentDir, workerDir string) error {
	extensionsDir := filepath.Join(workerDir, "extensions")
	if info, err := os.Lstat(extensionsDir); err == nil && info.Mode()&os.ModeSymlink != 0 {
		if err := os.Remove(extensionsDir); err != nil {
			return err
		}
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.MkdirAll(extensionsDir, 0700); err != nil {
		return err
	}
	sourceDir := filepath.Join(agentDir, "extensions")
	entries, err := os.ReadDir(sourceDir)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, entry := range entries {
		if entry.Name() == localCompactionGuardFilename {
			continue
		}
		target := filepath.Join(extensionsDir, entry.Name())
		if _, err := os.Lstat(target); err == nil {
			continue
		}
		if err := os.Symlink(filepath.Join(sourceDir, entry.Name()), target); err != nil && !os.IsExist(err) {
			return err
		}
	}
	return os.WriteFile(filepath.Join(extensionsDir, localCompactionGuardFilename), []byte(localCompactionGuardSource), 0600)
}

func validateProjectCompactionSettings(sessionPath string, minimumReserve, maximumKeepRecent int) error {
	file, err := os.Open(sessionPath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		_ = file.Close()
		if err := scanner.Err(); err != nil {
			return err
		}
		return fmt.Errorf("Local Mode session has no header")
	}
	_ = file.Close()
	var header struct {
		CWD string `json:"cwd"`
	}
	if err := json.Unmarshal(scanner.Bytes(), &header); err != nil {
		return fmt.Errorf("parse Local Mode session header: %w", err)
	}
	if header.CWD == "" {
		return fmt.Errorf("Local Mode session header has no working directory")
	}
	data, err := os.ReadFile(filepath.Join(header.CWD, ".pi", "settings.json"))
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	var projectSettings map[string]any
	if err := json.Unmarshal(data, &projectSettings); err != nil {
		return fmt.Errorf("parse project .pi/settings.json for Local Mode: %w", err)
	}
	projectCompaction, _ := projectSettings["compaction"].(map[string]any)
	if projectCompaction == nil {
		return nil
	}
	if enabled, ok := projectCompaction["enabled"].(bool); ok && !enabled {
		return fmt.Errorf("project .pi/settings.json disables compaction; enable it before using Local Mode")
	}
	if reserve, ok := numberValue(projectCompaction["reserveTokens"]); ok && reserve < minimumReserve {
		return fmt.Errorf("project .pi/settings.json compaction.reserveTokens must be at least %d for Local Mode", minimumReserve)
	}
	if keep, ok := numberValue(projectCompaction["keepRecentTokens"]); ok && keep > maximumKeepRecent {
		return fmt.Errorf("project .pi/settings.json compaction.keepRecentTokens must be at most %d for Local Mode", maximumKeepRecent)
	}
	return nil
}

func numberValue(value any) (int, bool) {
	switch number := value.(type) {
	case float64:
		return int(number), true
	case int:
		return number, true
	case json.Number:
		parsed, err := number.Int64()
		return int(parsed), err == nil
	default:
		return 0, false
	}
}
