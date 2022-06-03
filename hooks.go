package main

import (
	"os"
	"os/exec"
)

const commitHooksFilename = "./hooks.sema"

func hooks() {
	if commitHooksFileExists() {
		prepareCommand().Run()
	}
}

func commitHooksFileExists() bool {
	_, err := os.Open(commitHooksFilename)
	return err == nil
}

func prepareCommand() (cmd *exec.Cmd) {
	cmd = exec.Command(commitHooksFilename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return
}
