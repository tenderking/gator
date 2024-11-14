package main

import (
	"context"
	"fmt"
	"gator/internal/config"
	"gator/internal/rss"
	"time"
)

func handlerAgg(s *config.State, cmd config.Command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time>")
	}

	time_btw_reqs := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(time_btw_reqs)
	if err != nil {
		return fmt.Errorf("error parsing time: %w. Please provide a valid duration string like 1s, 1m, 1h, etc.", err)
	}

	fmt.Println("Collecting feeds every", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *config.State) error {
	ctx := context.Background()
	feed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}
	err = s.Db.MarkFeedAsFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("error setting feed last fetched: %w", err)
	}
	f, err := rss.FetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	for _, item := range f.Channel.Item {
		fmt.Printf("Title: %+v\n", item.Title)
	}

	return nil
}
