package main

type pipeline struct {
	commands []func()
}

func pipe(commands ...func()) *pipeline {
	return &pipeline{
		commands: commands,
	}
}

func (p *pipeline) then(commands ...func()) *pipeline {
	p.commands = append(p.commands, commands...)
	return p
}

func (p *pipeline) thenIf(condition bool, commands ...func()) *pipeline {
	if condition {
		p.then(commands...)
	}
	return p
}

func (p *pipeline) run() {
	for _, command := range p.commands {
		command()
	}
}
