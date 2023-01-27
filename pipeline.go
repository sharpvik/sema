package main

type pipeline struct {
	commands []command
}

type command func() error

func pipe(commands ...command) *pipeline {
	return &pipeline{
		commands: commands,
	}
}

func (p *pipeline) then(commands ...command) *pipeline {
	p.commands = append(p.commands, commands...)
	return p
}

func (p *pipeline) thenIf(condition bool, commands ...command) *pipeline {
	if condition {
		p.then(commands...)
	}
	return p
}

func (p *pipeline) run() error {
	for _, command := range p.commands {
		if err := command(); err != nil {
			return err
		}
	}
	return nil
}
