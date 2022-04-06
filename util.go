package main

import (
	"os"
	"os/exec"
)

func abort(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func runCommandAndAbortOnError(command *exec.Cmd) {
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	abort(command.Run())
}
