package main

import (
	"fmt"
	"gator/internal/config"
	"log"
	"os"
)

func main() {
	// Create a new instance of the server

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	configState := &config.State{
		Config: &cfg,
	}
	// Initialize commands and register handlers
	commands := &config.Commands{}
	commands.Register("login", config.HandlerLogin)
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

	fmt.Println(cfg)

}
