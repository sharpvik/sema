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
		Long bool
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

func (r *Agent) Hooks() (err error) {
	if commitHooksFileExists() {
		return try(exec.Command(commitHooksFilename))
	}
	return
}

func (r *Agent) Title() (_ error) {
	r.commitTitle = fmt.Sprintf("%s(%s): %s", label(), scope(), synopsis())
	return
}

func (r *Agent) Add() (err error) {
	return try(exec.Command("git", "add", "."))
}

func (r *Agent) Commit() (err error) {
	if r.Config.Commit.Long {
		return r.longCommit()
	} else {
		return r.shortCommit()
	}
}

func (r *Agent) Push() (err error) {
	args := []string{"push"}
	if r.Config.Push.Force {
		args = append(args, "-f")
	}
	return try(exec.Command("git", args...))
}

func (r *Agent) longCommit() (err error) {
	path, err := r.createCommitFile()
	if err != nil {
		return
	}
	return try(exec.Command("git", "commit", "-t", path))
}

func (r *Agent) createCommitFile() (path string, err error) {
	file, err := os.CreateTemp("", "sema-commit-template-")
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(r.commitTitle + "\n\n")
	return file.Name(), err
}

func (r *Agent) shortCommit() (err error) {
	display(r.commitTitle)
	return try(exec.Command("git", "commit", "-m", r.commitTitle))
}
