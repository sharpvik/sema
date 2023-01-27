package labels

import (
	"fmt"
)

var list = [9]Label{
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
	{
		Name:   "infra",
		Reason: "Infrastructural changes",
	},
}

type Label struct {
	Name   string
	Reason string
}

func (label *Label) String() string {
	return fmt.Sprintf("%-10s%s", label.Name, label.Reason)
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
