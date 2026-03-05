package main

import (
	"os/exec"
	"testing"
)

func TestCLIOutput(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("failed to run CLI: %v", err)
	}

	want := "Hello, Go!\n"
	if got := string(out); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
