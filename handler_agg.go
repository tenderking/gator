package main

import (
	"context"
	"fmt"
	"gator/internal/config"
	"gator/internal/rss"
	"log"
)

func handlerAgg(s *config.State, cmd config.Command) error {
	if len(cmd.Args) != 0 {
		log.Fatalf("usage: agg <feedURL>")
	}
	const feedURL = "https://www.wagslane.dev/index.xml"

	ctx := context.Background()
	feed, err := rss.FetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	fmt.Println(feed.Channel)

	return nil
}
