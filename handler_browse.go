package main

import (
	"context"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"strconv"

	"github.com/google/uuid"
)

func handlerBrowsePosts(s *config.State, cmd config.Command, user database.User) error {

	var limit int
	switch len(cmd.Args) {
	case 0:
		limit = 2
	case 1:
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}
	default:
		return fmt.Errorf("usage: browse <limit>")
	}

	if limit < 1 {
		return fmt.Errorf("limit must be a positive integer")
	}
	ctx := context.Background()
	params := database.GetUserPostsParams{
		UserID: uuid.NullUUID{UUID: user.ID, Valid: true},
		Limit:  int32(limit),
	}
	posts, err := s.Db.GetUserPosts(ctx, params)
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}
	fmt.Println("Posts:", len(posts))
	for _, post := range posts {
		fmt.Println("Title: ", post.Title)
		fmt.Println("Url: ", post.Url)
		fmt.Println("Description: ", post.Description)
		fmt.Println()
	}
	return nil
}
