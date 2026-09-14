package rpc

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeLocalModeTestSession(t *testing.T, root string) (string, string) {
	t.Helper()
	projectDir := filepath.Join(root, "project")
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		t.Fatal(err)
	}
	sessionPath := filepath.Join(root, "session.jsonl")
	header, _ := json.Marshal(map[string]any{"type": "session", "id": "session", "cwd": projectDir})
	if err := os.WriteFile(sessionPath, append(header, '\n'), 0600); err != nil {
		t.Fatal(err)
	}
	return sessionPath, projectDir
}

func TestPrepareLocalWorkerAgentDirSetsExactSixtyFivePercentThreshold(t *testing.T) {
	root := t.TempDir()
	agentDir := filepath.Join(root, "agent")
	if err := os.MkdirAll(agentDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(agentDir, "settings.json"), []byte(`{"theme":"dark","compaction":{"keepRecentTokens":12000}}`), 0600); err != nil {
		t.Fatal(err)
	}
	sessionPath, _ := writeLocalModeTestSession(t, root)
	workerDir, err := prepareLocalWorkerAgentDir(agentDir, sessionPath, 100000)
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(workerDir, "settings.json"))
	if err != nil {
		t.Fatal(err)
	}
	var settings map[string]any
	if err := json.Unmarshal(data, &settings); err != nil {
		t.Fatal(err)
	}
	compaction := settings["compaction"].(map[string]any)
	if got := int(compaction["reserveTokens"].(float64)); got != 35001 {
		t.Fatalf("reserveTokens = %d, want 35001", got)
	}
	if got := int(compaction["keepRecentTokens"].(float64)); got != 12000 {
		t.Fatalf("keepRecentTokens = %d, want preserved stricter value 12000", got)
	}
	if settings["theme"] != "dark" {
		t.Fatalf("unrelated global setting was not preserved: %#v", settings)
	}
}

func TestPrepareLocalWorkerAgentDirInstallsCompactionGuardAndPreservesExtensions(t *testing.T) {
	root := t.TempDir()
	agentDir := filepath.Join(root, "agent")
	extensionsDir := filepath.Join(agentDir, "extensions")
	if err := os.MkdirAll(extensionsDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(extensionsDir, "custom.js"), []byte("export default function () {}\n"), 0600); err != nil {
		t.Fatal(err)
	}
	sessionPath, _ := writeLocalModeTestSession(t, root)
	legacyWorkerDir := filepath.Join(agentDir, "pi-web", "local-workers", "session")
	if err := os.MkdirAll(legacyWorkerDir, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(extensionsDir, filepath.Join(legacyWorkerDir, "extensions")); err != nil {
		t.Fatal(err)
	}
	workerDir, err := prepareLocalWorkerAgentDir(agentDir, sessionPath, 100000)
	if err != nil {
		t.Fatal(err)
	}
	customInfo, err := os.Lstat(filepath.Join(workerDir, "extensions", "custom.js"))
	if err != nil {
		t.Fatal(err)
	}
	if customInfo.Mode()&os.ModeSymlink == 0 {
		t.Fatal("existing user extension was not preserved as a symlink")
	}
	workerExtensionsInfo, err := os.Lstat(filepath.Join(workerDir, "extensions"))
	if err != nil {
		t.Fatal(err)
	}
	if workerExtensionsInfo.Mode()&os.ModeSymlink != 0 {
		t.Fatal("legacy extensions symlink was not replaced with an isolated directory")
	}
	guard, err := os.ReadFile(filepath.Join(workerDir, "extensions", localCompactionGuardFilename))
	if err != nil {
		t.Fatal(err)
	}
	for _, required := range []string{"session_compact_failed", "session_compact", "requiredReduction", "ctx.abort()"} {
		if !strings.Contains(string(guard), required) {
			t.Fatalf("compaction guard does not contain %q", required)
		}
	}
	if _, err := os.Stat(filepath.Join(extensionsDir, localCompactionGuardFilename)); !os.IsNotExist(err) {
		t.Fatalf("guard leaked into the user's extension directory: %v", err)
	}
}

func TestPrepareLocalWorkerAgentDirRejectsProjectOverrideThatWeakensThreshold(t *testing.T) {
	root := t.TempDir()
	agentDir := filepath.Join(root, "agent")
	if err := os.MkdirAll(agentDir, 0755); err != nil {
		t.Fatal(err)
	}
	sessionPath, projectDir := writeLocalModeTestSession(t, root)
	settingsDir := filepath.Join(projectDir, ".pi")
	if err := os.MkdirAll(settingsDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(settingsDir, "settings.json"), []byte(`{"compaction":{"reserveTokens":1000}}`), 0600); err != nil {
		t.Fatal(err)
	}
	_, err := prepareLocalWorkerAgentDir(agentDir, sessionPath, 100000)
	if err == nil || !strings.Contains(err.Error(), "at least 35001") {
		t.Fatalf("error = %v", err)
	}
}
