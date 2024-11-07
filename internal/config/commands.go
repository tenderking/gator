package config

import (
	"log"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	registeredCommands map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if c.registeredCommands == nil {
		c.registeredCommands = make(map[string]func(*State, Command) error)
	}
	c.registeredCommands[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.registeredCommands[cmd.Name]
	if !ok {
		log.Fatalf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
