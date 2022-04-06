package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	. "github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
)

func more() {
	var builder strings.Builder
	builder.WriteString("Labels explained:\n\n")
	for _, label := range labels {
		builder.WriteString("    " + label.String() + "\n")
	}
	fmt.Println(builder.String())
}

func add() {
	runCommandAndAbortOnError(exec.Command("git", "add", "."))
}

func commit() {
	label := label()
	scope := scope()
	synopsis := message()
	message := fmt.Sprintf("%s(%s): %s", label, scope, synopsis)
	display(message)
	runCommandAndAbortOnError(exec.Command("git", "commit", "-m", message))
}

func push() {
	runCommandAndAbortOnError(exec.Command("git", "push"))
}

func label() (choice string) {
	prompt := promptui.Select{
		Label: "Select commit label",
		Items: tagsOnly(),
	}
	_, choice, err := prompt.Run()
	abort(err)
	return
}

func scope() (scope string) {
	valiadtor := func(input string) (err error) {
		if len(input) > 15 {
			return errors.New("input too long")
		}
		return
	}
	prompt := promptui.Prompt{
		Label:    "Change scope",
		Validate: valiadtor,
	}
	scope, err := prompt.Run()
	abort(err)
	return
}

func message() (msg string) {
	prompt := promptui.Prompt{Label: "Commit message"}
	msg, err := prompt.Run()
	abort(err)
	return
}

func display(message string) {
	fmt.Printf("Commit: %v\n\n", Green(message))
}
