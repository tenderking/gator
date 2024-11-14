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

func handlerAddFeed(s *config.State, cmd config.Command, user database.User) error {
	if len(cmd.Args) != 2 {
		log.Fatalf("usage: addfeed <feedName> <feedurl> ")
		return nil
	}

	ctx := context.Background()

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	feed, err := s.Db.CreateUserFeed(ctx, database.CreateUserFeedParams{
		ID:        uuid.New(),
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	s.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
		FeedID:    uuid.NullUUID{UUID: feed.ID, Valid: true},
		UpdatedAt: time.Now(),
	})

	return nil
}
