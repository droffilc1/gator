package main

import (
	"fmt"
	"log"
	"os"

	"github.com/droffilc1/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := &state{
		config: &cfg,
	}

	cmds := &commands{
		commands: map[string]func(*state, command) error{
			"login": handlerLogin,
		},
	}

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments were provided")
		os.Exit(1)
	}

	c := command{
		name:     os.Args[1],
		username: os.Args[2:],
	}

	err = cmds.run(s, c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read again from disk to confirm write worked
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", cfg)
}