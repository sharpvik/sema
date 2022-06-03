package main

import "os"

func init() {
	parseFlags()
	if meta() {
		os.Exit(0)
	}
}

func main() {
	hooks()
	if *flags.add {
		add()
	}
	commit()
	if *flags.push {
		push()
	}
}
