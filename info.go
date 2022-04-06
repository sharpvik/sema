package main

import "fmt"

const (
	name       = "sema"
	version    = "v0.1.2"
	maintainer = "Viktor A. Rozenko Voitenko"
	email      = "sharp.vik@gmail.com"
)

func info() {
	fmt.Printf("%s %s by %s <%s>\n\n", name, version, maintainer, email)
}
