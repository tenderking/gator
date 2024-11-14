package main

import (
	"context"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *config.State, cmd config.Command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <feed-url>")
	}

	ctx := context.Background()

	feedURL := cmd.Args[0]

	feed, err := s.Db.GetFeedByURL(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	feedFollow, err := s.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
		FeedID:    uuid.NullUUID{UUID: feed.ID, Valid: true},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("Followed feed: %s\n", feedFollow.FeedName)
	return nil
}

func handlerUnfollow(s *config.State, cmd config.Command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: unfollow <feed-url>")
	}

	feedURL := cmd.Args[0]
	ctx := context.Background()
	err := s.Db.UnfollowFeed(ctx, database.UnfollowFeedParams{
		Url: feedURL,
		ID:  user.ID,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}

	fmt.Printf("Unfollowed feed: %s\n", feedURL)

	return nil
}

func handlerFollowing(s *config.State, cmd config.Command, user database.User) error {

	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: following")
	}

	ctx := context.Background()

	follows, err := s.Db.GetFeedFollowsForUser(ctx, uuid.NullUUID{UUID: user.ID, Valid: true})
	if err != nil {
		return fmt.Errorf("error getting follows: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Println("Following:")
	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}
