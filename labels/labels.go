package labels

import (
	"fmt"
)

type Label struct {
	Name   string
	Reason string
}

func (label *Label) String() string {
	return fmt.Sprintf("%-10s%s", label.Name, label.Reason)
}

var list = [8]Label{
	{
		Name:   "feat",
		Reason: "New feature",
	},
	{
		Name:   "fix",
		Reason: "Bug fix",
	},
	{
		Name:   "docs",
		Reason: "Documentation improvement",
	},
	{
		Name:   "style",
		Reason: "Code formatting",
	},
	{
		Name:   "refactor",
		Reason: "Code refactoring",
	},
	{
		Name:   "test",
		Reason: "Test suite improvement",
	},
	{
		Name:   "perf",
		Reason: "Performance improvement",
	},
	{
		Name:   "chore",
		Reason: "Grunt task",
	},
}

func Explained() (labels []string) {
	labels = make([]string, len(list))
	for i, label := range list {
		labels[i] = label.String()
	}
	return
}

func Get(index int) Label {
	return list[index]
}
