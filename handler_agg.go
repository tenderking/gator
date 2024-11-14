package main

import (
	"context"
	"fmt"
	"gator/internal/config"
	"gator/internal/rss"
)

func handlerAgg(s *config.State, cmd config.Command) error {

	const feedURL = "https://www.wagslane.dev/index.xml"
	ctx := context.Background()

	feed, err := rss.FetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed.Channel)

	return nil
}
