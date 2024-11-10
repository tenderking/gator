package main

import (
	"context"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *config.State, cmd config.Command) error {

	if len(cmd.Args) != 1 {
		log.Fatalf("usage: register <username> ")
	}

	username := cmd.Args[0]
	ctx := context.Background()
	user, err := s.Db.CreateUser(ctx, database.CreateUserParams{

		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	fmt.Println("User", username, "created")
	s.Config.CurrentUserName = user.Name
	err = config.SetUser(s.Config, user.Name)
	if err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}
	fmt.Println("User created successfully:")
	printUser(user)

	return nil
}
func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
