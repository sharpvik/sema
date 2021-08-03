package main

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"os"
	"os/exec"
)

func abort(err error) {
	if err != nil {
		fmt.Println(Red(err.Error()))
		os.Exit(1)
	}
}

func runCommandAndAbortOnError(command *exec.Cmd) {
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	abort(command.Run())
}
