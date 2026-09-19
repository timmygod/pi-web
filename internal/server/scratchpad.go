package server

import (
	"database/sql"
	"net/http"
	"strings"
	"time"
)

// lookupScratchpad returns the saved scratchpad content for a project path.
// An unknown project (or no database) yields an empty string, not an error, so
// callers on the page-render path can pre-fill the textarea best-effort.
func (s *Server) lookupScratchpad(project string) (string, error) {
	if project == "" || s.db == nil {
		return "", nil
	}
	var content string
	err := s.db.QueryRow("SELECT content FROM scratchpads WHERE project_path = ?", project).Scan(&content)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return content, err
}

func (s *Server) handleGetScratchpad(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	project := r.URL.Query().Get("project")
	if project == "" {
		writeJSONError(w, http.StatusBadRequest, "project query parameter is required")
		return
	}

	if s.db == nil {
		writeJSONError(w, http.StatusInternalServerError, "database is unavailable")
		return
	}

	content, err := s.lookupScratchpad(project)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to query scratchpad: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{"content": content})
}

func (s *Server) handleSaveScratchpad(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var body struct {
		Project string `json:"project"`
		Content string `json:"content"`
		Mode    string `json:"mode"`
	}
	if !decodeJSONBody(w, r, &body) {
		return
	}

	if body.Project == "" {
		writeJSONError(w, http.StatusBadRequest, "project is required")
		return
	}

	mode := strings.TrimSpace(strings.ToLower(body.Mode))
	if mode == "" {
		mode = "replace"
	}
	if mode != "replace" && mode != "append" {
		writeJSONError(w, http.StatusBadRequest, "mode must be replace or append")
		return
	}

	if s.db == nil {
		writeJSONError(w, http.StatusInternalServerError, "database is unavailable")
		return
	}

	var err error
	if mode == "append" {
		_, err = s.db.Exec(`INSERT INTO scratchpads (project_path, content, updated_at)
			VALUES (?, ?, ?)
			ON CONFLICT(project_path) DO UPDATE SET
				content = scratchpads.content || excluded.content,
				updated_at = excluded.updated_at`,
			body.Project, body.Content, time.Now())
	} else {
		_, err = s.db.Exec(`INSERT INTO scratchpads (project_path, content, updated_at)
			VALUES (?, ?, ?)
			ON CONFLICT(project_path) DO UPDATE SET content=excluded.content, updated_at=excluded.updated_at`,
			body.Project, body.Content, time.Now())
	}
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to save scratchpad: "+err.Error())
		return
	}

	content := body.Content
	if mode == "append" {
		if stored, lookupErr := s.lookupScratchpad(body.Project); lookupErr == nil {
			content = stored
		}
	}
	s.notifyScratchpadChanged(body.Project, content)
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "content": content})
}

func (s *Server) notifyScratchpadChanged(project, content string) {
	payload := map[string]any{"project": project, "content": content}
	if msg, err := formatSSEJSONEvent("scratchpad", payload); err == nil {
		s.broadcast(globalSessID, msg)
	}
}
