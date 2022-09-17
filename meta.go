package main

import (
	"fmt"

	"github.com/pborman/getopt/v2"
)

const (
	name       = "sema"
	version    = "v2.1.0"
	maintainer = "Viktor A. Rozenko Voitenko"
	email      = "sharp.vik@gmail.com"
	url        = "https://github.com/sharpvik/sema"
)

func info() (_ error) {
	fmt.Printf("%s %s by %s <%s>\n  💾 GitHub %s\n",
		name, version, maintainer, email, url)
	return
}

func usage() (_ error) {
	getopt.Usage()
	return
}

func newline() (_ error) {
	fmt.Println()
	return
}
