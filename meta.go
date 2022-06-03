package main

import (
	"fmt"
	"strings"
)

const (
	name       = "sema"
	version    = "v0.2.3"
	maintainer = "Viktor A. Rozenko Voitenko"
	email      = "sharp.vik@gmail.com"
)

func info() {
	fmt.Printf("%s %s by %s <%s>\n\n", name, version, maintainer, email)
}

func more() {
	var builder strings.Builder
	builder.WriteString("Labels explained:\n\n")
	for _, label := range labels {
		builder.WriteString("    " + label.String() + "\n")
	}
	fmt.Println(builder.String())
}
