package main

import (
	"github.com/pborman/getopt/v2"
	"github.com/sharpvik/sema/agent"
	"github.com/sharpvik/sema/labels"
)

type flags struct {
	/* Operational flags that affect functional behaviour. */
	add   *bool
	push  *bool
	force *bool
	long  *bool

	/* Meta flags that display information about the program. */
	version *bool
	help    *bool
	more    *bool
}

func (f *flags) parse() *flags {
	f.add = getopt.BoolLong("add", 'a', "Begin with running 'git add'")
	f.push = getopt.BoolLong("push", 'p', "Run 'git push' on successful commit")
	f.force = getopt.BoolLong("force", 'f', "Add force push flag '-f' during 'git push'")
	f.long = getopt.BoolLong("long", 'l', "Open editor to elaborate")

	f.version = getopt.BoolLong("version", 'n', "Display installed version of sema")
	f.help = getopt.BoolLong("help", 'h', "Display help message")
	f.more = getopt.BoolLong("more", 'm', "Explain commit types")

	getopt.Parse()
	return f
}

func (f *flags) runner() *pipeline {
	if *f.version {
		return pipe(info)
	} else if *f.help {
		return pipe(info, getopt.Usage)
	} else if *f.more {
		return pipe(labels.Explain)
	}
	return f.sema()
}

func (f *flags) sema() *pipeline {
	run := agent.New(f.config())
	return pipe(run.Title, run.Hooks).
		thenIf(*f.add, run.Add).
		then(run.Commit).
		thenIf(*f.push, run.Push)
}

func (f *flags) config() *agent.Config {
	return &agent.Config{
		Commit: agent.Commit{Long: *f.long},
		Push:   agent.Push{Force: *f.force},
	}
}
