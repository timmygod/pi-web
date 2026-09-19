package rpc

import (
	"os"
	"testing"
)

func TestDetachedPiDirIsNotEmpty(t *testing.T) {
	dir := detachedPiDir()
	if dir == "" {
		t.Fatal("detachedPiDir must be a real directory so pi does not inherit pi-web's checkout")
	}
	if info, err := os.Stat(dir); err != nil || !info.IsDir() {
		t.Fatalf("detachedPiDir = %q, stat err %v", dir, err)
	}
}
