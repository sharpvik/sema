package main

import "fmt"

type Label struct {
	tag         string
	description string
}

func (label *Label) String() string {
	return fmt.Sprintf("%-10s%s", label.tag, label.description)
}

var labels = [8]Label{
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

func tagsOnly() (tags []string) {
	tags = make([]string, len(labels))
	for i, label := range labels {
		tags[i] = label.tag
	}
	return
}
