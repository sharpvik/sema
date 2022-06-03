package main

import (
	"os"
	"os/exec"
)

const commitHooksFilename = "./hooks.sema"

func hooks() {
	if commitHooksFileExists() {
		runCommandAndAbortOnError(exec.Command(commitHooksFilename))
	}
}

func commitHooksFileExists() bool {
	file, err := os.Open(commitHooksFilename)
	defer file.Close()
	return err == nil
}
