package main

import (
	"context"
	"fmt"
	"gator/internal/config"
)

func handlerGetUsers(s *config.State, cmd config.Command) error {
	ctx := context.Background()
	users, err := s.Db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)
	}
	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Print(user.Name, " (current)")
		} else {
			fmt.Println(user.Name)
		}
	}
	return nil
}
