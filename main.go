package main

import (
	"database/sql"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Create a new instance of the server

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	configState := &config.State{
		Db:     dbQueries,
		Config: &cfg,
	}
	// Initialize commands and register handlers
	commands := &config.Commands{}
	commands.Register("login", HandlerLogin)
	commands.Register("register", handlerRegister)
	commands.Register("reset", handlerReset)
	commands.Register("users", handlerGetUsers)
	commands.Register("agg", handlerAgg)
	commands.Register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.Register("feeds", handlerGetFeeds)
	commands.Register("follow", middlewareLoggedIn(handlerFollow))
	commands.Register("following", middlewareLoggedIn(handlerFollowing))
	commands.Register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.Register("browse", middlewareLoggedIn(handlerBrowsePosts))

	if len(os.Args) < 2 {
		log.Fatalf("usage: gator <command> [<args>]")
	}
	commandName := os.Args[1]
	commandArgs := os.Args[2:]
	cmd := config.Command{
		Name: commandName,
		Args: commandArgs,
	}
	err = commands.Run(configState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
