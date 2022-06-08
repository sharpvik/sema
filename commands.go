package main

import (
	"errors"
	"fmt"
	"os/exec"

	. "github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
)

func commitMessage() string {
	return fmt.Sprintf("%s(%s): %s", label(), scope(), synopsis())
}

func add() {
	abortOnError(exec.Command("git", "add", "."))
}

func commit(message string) {
	display(message)
	abortOnError(exec.Command("git", "commit", "-m", message))
}

func push() {
	args := []string{"push"}
	if *flags.force {
		args = append(args, "-f")
	}
	abortOnError(exec.Command("git", args...))
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

func synopsis() (message string) {
	prompt := promptui.Prompt{Label: "Commit message"}
	message, err := prompt.Run()
	abort(err)
	return
}

func display(message string) {
	fmt.Printf("Commit: %v\n\n", Green(message))
}
