package agent

import (
	"fmt"
	"os"
	"os/exec"
)

type (
	Agent struct {
		Config      *Config
		commitTitle string
	}

	Config struct {
		Commit Commit
		Push   Push
	}

	Commit struct {
		Long     bool
		Breaking bool
	}

	Push struct {
		Force bool
	}
)

const commitHooksFilename = "./hooks.sema"

func New(config *Config) *Agent {
	return &Agent{
		Config: config,
	}
}

func (a *Agent) Hooks() (err error) {
	if !commitHooksFileExists() {
		return
	}
	if err := try(exec.Command(commitHooksFilename)); err != nil {
		return fmt.Errorf("Commit hooks failed: %s", err)
	}
	return
}

func (a *Agent) Title() (_ error) {
	a.commitTitle = fmt.Sprintf("%s%s: %s", a.label(), scope(), synopsis())
	return
}

func (a *Agent) Add() (err error) {
	return try(exec.Command("git", "add", "."))
}

func (a *Agent) Commit() (err error) {
	if a.Config.Commit.Long {
		return a.longCommit()
	} else {
		return a.shortCommit()
	}
}

func (a *Agent) Push() (err error) {
	args := []string{"push"}
	if a.Config.Push.Force {
		args = append(args, "-f")
	}
	return try(exec.Command("git", args...))
}

func (a *Agent) longCommit() (err error) {
	path, err := a.createCommitFile()
	if err != nil {
		return
	}
	return try(exec.Command("git", "commit", "-t", path))
}

func (a *Agent) createCommitFile() (path string, err error) {
	file, err := os.CreateTemp("", "sema-commit-template-")
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(a.commitTitle + "\n\n" + a.maybeBreakingSuffix())
	return file.Name(), err
}

func (a *Agent) shortCommit() (err error) {
	display(a.commitTitle)
	return try(exec.Command("git", "commit", "-m", a.commitTitle))
}
