package labels

import (
	"fmt"
	"strings"
)

type Label struct {
	tag         string
	description string
}

func (label *Label) String() string {
	return fmt.Sprintf("%-10s%s", label.tag, label.description)
}

var list = [8]Label{
	{
		tag:         "feat",
		description: "new feature for the user",
	},
	{
		tag:         "fix",
		description: "bug fix for the user",
	},
	{
		tag:         "docs",
		description: "changes to the documentation",
	},
	{
		tag:         "style",
		description: "formatting with no production code change",
	},
	{
		tag:         "refactor",
		description: "refactoring production code",
	},
	{
		tag:         "test",
		description: "adding missing tests, refactoring tests",
	},
	{
		tag:         "perf",
		description: "performance improvements",
	},
	{
		tag:         "chore",
		description: "updating grunt tasks",
	},
}

func TagsOnly() (tags []string) {
	tags = make([]string, len(list))
	for i, label := range list {
		tags[i] = label.tag
	}
	return
}

func Explain() {
	var builder strings.Builder
	builder.WriteString("Labels explained:\n\n")
	for _, label := range list {
		builder.WriteString("    " + label.String() + "\n")
	}
	fmt.Println(builder.String())
}
