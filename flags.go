package main

import "github.com/pborman/getopt/v2"

type flagsStore struct {
	/* Operational flags that affect functional behaviour. */
	add  *bool
	push *bool

	/* Meta flags that display information about the program. */
	help    *bool
	more    *bool
	version *bool
}

var flags flagsStore

func parseFlags() {
	flags.help = getopt.BoolLong("help", 'h', "Display help message")
	flags.more = getopt.BoolLong("more", 'm', "Explain commit types")
	flags.add = getopt.BoolLong("add", 'a', "Begin with running 'git add'")
	flags.push = getopt.BoolLong("push", 'p', "Run 'git push' on successful commit")
	flags.version = getopt.BoolLong("version", 'n', "Display installed version of sema")

	getopt.Parse()
}

func meta() (exit bool) {
	if *flags.help {
		info()
		getopt.Usage()
		exit = true
	} else if *flags.more {
		more()
		exit = true
	} else if *flags.version {
		info()
		exit = true
	}
	return
}
