package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registerCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.registerCommands[cmd.Name]

	if !ok {
		return fmt.Errorf("unknown command: %s\n", cmd.Name)
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registerCommands[name] = f
}