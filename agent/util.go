package agent

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/sharpvik/sema/labels"
)

func label() (choice string) {
	prompt := promptui.Select{
		Label: "Select commit label",
		Items: labels.Explained(),
	}
	i, _, err := prompt.Run()
	AbortIfError(err)
	return labels.Get(i).Name
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
	fmt.Printf("Commit: %v\n\n", Green(message))
}

func commitHooksFileExists() bool {
	file, err := os.Open(commitHooksFilename)
	defer file.Close()
	return err == nil
}

func AbortIfError(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func try(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func bracketedOrEmpty(label string) (l string) {
	if label == "" {
		return
	}
	return "(" + label + ")"
}
