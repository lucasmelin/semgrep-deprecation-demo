package main

import (
	"os"
	runner "os/exec"
	"strings"
)

func main() {
	program := "demo"
	if len(os.Args) >= 2 {
		program = os.Args[1]
	}
	runApp(program)
}

func runApp(program string) string {
	cmd := runner.Command("code-buddy", "run", program, "--please")
	stdout := &strings.Builder{}
	cmd.Stdout = stdout
	_ = cmd.Run()
	return stdout.String()
}
