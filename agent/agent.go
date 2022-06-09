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

func (r *Agent) Hooks() {
	if commitHooksFileExists() {
		abortOnError(exec.Command(commitHooksFilename))
	}
}

func (r *Agent) Title() {
	r.commitTitle = fmt.Sprintf("%s(%s): %s", label(), scope(), synopsis())
}

func (r *Agent) Add() {
	abortOnError(exec.Command("git", "add", "."))
}

func (r *Agent) Commit() {
	if r.Config.Commit.Long {
		r.longCommit()
	} else {
		r.shortCommit()
	}
}

func (r *Agent) Push() {
	args := []string{"push"}
	if r.Config.Push.Force {
		args = append(args, "-f")
	}
	abortOnError(exec.Command("git", args...))
}

func (r *Agent) longCommit() {
	path, err := r.createCommitFile()
	abort(err)
	abortOnError(exec.Command("git", "commit", "-t", path))
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

func (r *Agent) shortCommit() {
	display(r.commitTitle)
	abortOnError(exec.Command("git", "commit", "-m", r.commitTitle))
}
