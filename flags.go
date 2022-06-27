package main

import (
	"io"

	"github.com/pborman/getopt/v2"
	"github.com/pkg/browser"
	"github.com/sharpvik/sema/agent"
)

type flags struct {
	/* Operational flags that affect functional behaviour. */
	add      *bool
	push     *bool
	force    *bool
	long     *bool
	breaking *bool

	/* Meta flags that display information about the program. */
	contribute *bool
	help       *bool
	version    *bool
}

func (f *flags) parse() *flags {
	f.add = getopt.BoolLong("add", 'a', "Begin by running `git add`")
	f.push = getopt.BoolLong("push", 'p', "Run `git push` on successful commit")
	f.force = getopt.BoolLong("force", 'f', "Force push changes with `git push -f`")
	f.long = getopt.BoolLong("long", 'l', "Open editor to elaborate commit message")
	f.breaking = getopt.BoolLong("breaking", 'b', "Mark commit as introducing breaking changes")

	f.version = getopt.BoolLong("version", 'v', "Display current version of sema")
	f.help = getopt.BoolLong("help", 'h', "Display help message")
	f.contribute = getopt.BoolLong("contribute", 'c', "Open sema GitHub repository in browser")

	getopt.Parse()
	return f
}

func (f *flags) runner() *pipeline {
	if *f.contribute {
		return pipe(openGitHub)
	} else if *f.help {
		return pipe(info, newline, usage)
	} else if *f.version {
		return pipe(info)
	}
	return f.sema()
}

func (f *flags) sema() *pipeline {
	run := agent.New(f.config())
	return pipe(run.Init, run.Title, run.Hooks).
		thenIf(*f.add, run.Add).
		then(run.Commit).
		thenIf(*f.push, run.Push)
}

func (f *flags) config() *agent.Config {
	return &agent.Config{
		Commit: agent.Commit{
			Long:     *f.long,
			Breaking: *f.breaking,
		},
		Push: agent.Push{Force: *f.force},
	}
}

func openGitHub() (_ error) {
	browser.Stdout = io.Discard
	browser.Stderr = io.Discard
	browser.OpenURL(url)
	return
}
