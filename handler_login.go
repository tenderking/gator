package main

import (
	"fmt"
	"gator/internal/config"
	"log"
)

func HandlerLogin(s *config.State, cmd config.Command) error {
	if len(cmd.Args) != 1 {
		log.Fatalf("usage: login <username>")
	}
	s.Config.CurrentUserName = cmd.Args[0]
	err := config.SetUser(s.Config, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}

	fmt.Println("You are now logged in as", cmd.Args[0])

	return nil
}
