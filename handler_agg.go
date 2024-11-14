package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"gator/internal/rss"
	"time"

	"github.com/google/uuid"
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

		err := s.Db.CreatePost(
			ctx,
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       item.Title,
				Url:         item.Link,
				Description: sql.NullString{String: item.Description, Valid: true},
				PublishedAt: func() time.Time {
					t, err := time.Parse(time.RFC1123, item.PubDate)
					if err != nil {
						return time.Time{}
					}
					return t
				}(),
				FeedID: uuid.NullUUID{UUID: feed.ID, Valid: true},
			},
		)
		if err != nil {
			return fmt.Errorf("error creating post: %w", err)
		}

	}

	return nil
}
