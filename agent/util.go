package agent

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/sharpvik/sema/labels"
)

func (a *Agent) label() (choice string) {
	prompt := promptui.Select{
		Label: "Select commit label",
		Items: labels.Explained(),
	}
	i, _, err := prompt.Run()
	AbortIfError(err)
	return labels.Get(i).Name + a.maybeBreakingExclam()
}

func scope() (scope string) {
	prompt := promptui.Prompt{Label: "Change scope"}
	scope, err := prompt.Run()
	AbortIfError(err)
	return bracketedOrEmpty(scope)
}

func synopsis() (message string) {
	prompt := promptui.Prompt{Label: "Commit message"}
	message, err := prompt.Run()
	AbortIfError(err)
	return
}

func display(message string) {
	fmt.Printf("Commit: %v\n\n", aurora.Green(message))
}

func AbortIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func try(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func bracketedOrEmpty(label string) string {
	if label == "" {
		return ""
	}
	return "(" + label + ")"
}

func (a *Agent) maybeBreakingExclam() string {
	if a.Config.Commit.Breaking {
		return "!"
	}
	return ""
}

func (a *Agent) maybeBreakingSuffix() string {
	if a.Config.Commit.Breaking {
		return "BREAKING CHANGE: \n"
	}
	return ""
}

func gitError(err error) error {
	if err == git.NoErrAlreadyUpToDate {
		return nil
	}
	return err
}
