package main

import "os"

func init() {
	parseFlags()
	if meta() {
		os.Exit(0)
	}
}

func main() {
	message := commitMessage()
	hooks()
	if *flags.add {
		add()
	}
	commit(message)
	if *flags.push {
		push()
	}
}
