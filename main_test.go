package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test_runApp(t *testing.T) {
	cmd := exec.Command("code-buddy", "run", "demo", "--please")
	stdout := &strings.Builder{}
	cmd.Stdout = stdout
	_ = cmd.Run()
	want := stdout.String()
	got := runApp("demo")
	if got != want {
		t.Fatalf("output does not match, wanted %q, got %q", want, got)
	}
}
