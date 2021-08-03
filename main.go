package main

import (
	"github.com/pborman/getopt/v2"
	"os"
)

type flagsStore struct {
	add  *bool
	push *bool
	help *bool
	more *bool
}

var flags = flagsStore{}

func init() {
	flags.help = getopt.BoolLong("help", 'h', "Display help message")
	flags.more = getopt.BoolLong("more", 'm', "Explain commit types")
	flags.add = getopt.BoolLong("add", 'a', "Begin with running 'git add'")
	flags.push = getopt.BoolLong(
		"push", 'p', "Run 'git push' on successful commit")

	getopt.Parse()

	if *flags.help {
		getopt.Usage()
		os.Exit(0)
	} else if *flags.more {
		more()
	}
}

func main() {
	if *flags.add {
		add()
	}
	commit()
	if *flags.push {
		push()
	}
}
