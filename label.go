package main

import "fmt"

type Label struct {
	tag         string
	description string
}

func (label *Label) String() string {
	return fmt.Sprintf("%-10s%s", label.tag, label.description)
}

var labels = [7]Label{
	{
		tag:         "Feat",
		description: "new feature for the user",
	},
	{
		tag:         "Fix",
		description: "bug fix for the user",
	},
	{
		tag:         "Docs",
		description: "changes to the documentation",
	},
	{
		tag:         "Style",
		description: "formatting with no production code change",
	},
	{
		tag:         "Refactor",
		description: "refactoring production code",
	},
	{
		tag:         "Test",
		description: "adding missing tests, refactoring tests",
	},
	{
		tag:         "Chore",
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
