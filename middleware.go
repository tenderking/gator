package main

import (
	"context"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
)

func middlewareLoggedIn(handler func(s *config.State, cmd config.Command, user database.User) error) func(*config.State, config.Command) error {
	return func(s *config.State, cmd config.Command) error {
		if s.Config.CurrentUserName == "" {
			return fmt.Errorf("not logged in")
		}
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error getting user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
