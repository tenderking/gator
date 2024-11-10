package main

import (
	"context"
	"fmt"
	"gator/internal/config"
	"log"
)

func handlerReset(s *config.State, cmd config.Command) error {
	if len(cmd.Args) != 0 {
		log.Fatalf("usage: reset")
	}
	ctx := context.Background()
	err := s.Db.DeleteAllUsers(ctx)
	if err != nil {
		return fmt.Errorf("error resetting database: %w", err)
	}
	fmt.Println("Database reset successfully")
	return nil
}
