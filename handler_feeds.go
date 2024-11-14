package main

import (
	"context"
	"fmt"
	"gator/internal/config"
)

func handlerGetFeeds(s *config.State, cmd config.Command) error {

	ctx := context.Background()

	feeds, err := s.Db.GetUserFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		name, err := s.Db.GetUserById(ctx, feed.UserID.UUID)
		if err != nil {
			return fmt.Errorf("error getting user: %w", err)
		}
		fmt.Println(name)

	}
	return nil
}
