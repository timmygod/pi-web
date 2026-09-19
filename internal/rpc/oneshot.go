package rpc

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// OneShot spawns `pi --mode rpc`, sends a single command, awaits the matching
// response, and tears the subprocess down. It exists so sessionless RPCs (e.g.
// get_available_models) don't reimplement spawn/scan/timeout machinery.
//
// The subprocess cwd is detached from pi-web's (see detachedPiDir). Extensions
// still load — including globally installed ones that register models — but
// project extensions from this checkout are not loaded on top of that copy.
func OneShot(ctx context.Context, command string, extraFields map[string]any) (json.RawMessage, error) {
	if _, err := exec.LookPath("pi"); err != nil {
		return nil, fmt.Errorf("pi executable not found: %w", err)
	}

	cmd := exec.CommandContext(ctx, "pi", "--mode", "rpc")
	cmd.Dir = detachedPiDir()
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	var stderrBuf strings.Builder
	cmd.Stderr = &stderrBuf

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	defer func() {
		_ = stdin.Close()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	}()

	reqID := fmt.Sprintf("oneshot-%d", time.Now().UnixNano())
	req := map[string]any{"id": reqID, "type": command}
	for k, v := range extraFields {
		req[k] = v
	}
	if err := WriteCommand(stdin, req); err != nil {
		return nil, err
	}

	type result struct {
		data json.RawMessage
		err  error
	}
	resCh := make(chan result, 1)
	go func() {
		scanner := bufio.NewScanner(stdout)
		scanner.Buffer(make([]byte, 0, 64*1024), 10*1024*1024)
		for scanner.Scan() {
			line := strings.TrimSuffix(scanner.Text(), "\r")
			if strings.TrimSpace(line) == "" {
				continue
			}
			var res response
			if err := json.Unmarshal([]byte(line), &res); err != nil {
				continue
			}
			if res.Type != "response" || res.ID != reqID {
				continue
			}
			if !res.Success {
				if res.Error != "" {
					resCh <- result{err: errors.New(res.Error)}
				} else {
					resCh <- result{err: fmt.Errorf("rpc %s rejected", command)}
				}
				return
			}
			resCh <- result{data: res.Data}
			return
		}
		if err := scanner.Err(); err != nil {
			resCh <- result{err: err}
			return
		}
		resCh <- result{err: fmt.Errorf("pi closed stdout without response (stderr: %q)", stderrBuf.String())}
	}()

	select {
	case r := <-resCh:
		return r.data, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
