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
		Items: labels.TagsOnly(),
	}
	_, choice, err := prompt.Run()
	abort(err)
	return
}

func scope() (scope string) {
	prompt := promptui.Prompt{Label: "Change scope"}
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

func commitHooksFileExists() bool {
	file, err := os.Open(commitHooksFilename)
	defer file.Close()
	return err == nil
}

func abort(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func abortOnError(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	abort(cmd.Run())
}
