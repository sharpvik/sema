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

func runCommandAndAbortOnError(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	abort(cmd.Run())
}
