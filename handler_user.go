package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/droffilc1/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.username) == 0 {
		return errors.New("username is required")
	}

	name := cmd.username[0]

	name, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("failed to get user: %s", err)
	}

	if err := s.config.SetUser(name); err != nil {
		return err
	}
	fmt.Println("user has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.username) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	name := cmd.username[0]
	now := time.Now()
	params, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %s", err)
	}

	if err := s.config.SetUser(params.Name); err != nil {
		return err
	}

	fmt.Println("User created successfully!")
	fmt.Printf(" * ID:        %v\n", params.ID)
	fmt.Printf(" * Name:      %v\n", params.Name)
	fmt.Printf(" * CreatedAt: %v\n", params.CreatedAt)
	return nil
}