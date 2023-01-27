package agent

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/sharpvik/sema/labels"
)

const defaultGitEditor = "vi"

func (a *Agent) label() string {
	prompt := promptui.Select{
		Label: "Select commit label",
		Items: labels.Explained(),
	}
	i, _, err := prompt.Run()
	abortIfError(err)
	return labels.Get(i).Name + a.maybeBreakingExclam()
}

func scope() string {
	prompt := promptui.Prompt{Label: "Change scope"}
	scope, err := prompt.Run()
	abortIfError(err)
	return bracketedOrEmpty(scope)
}

func synopsis() string {
	prompt := promptui.Prompt{Label: "Commit message"}
	message, err := prompt.Run()
	abortIfError(err)
	return message
}

func display(message string) {
	fmt.Printf("Commit: %v\n\n", aurora.Green(message))
}

func abortIfError(err error) {
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
