package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/droffilc1/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %s", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("username is required")
	}

	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("failed to get user: %s", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User switched successfully!")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list users: %s", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUsername {
			fmt.Printf(" * %v (current)\n", s.cfg.CurrentUsername)
			continue
		}
		fmt.Printf(" * %v\n", user.Name)
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:        %v\n", user.ID)
	fmt.Printf(" * Name:      %v\n", user.Name)
}