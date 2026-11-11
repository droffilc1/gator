package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.username) == 0 {
		return errors.New("username is required")
	}

	if err := s.config.SetUser(cmd.username[0]); err != nil {
		return err
	}
	fmt.Println("user has been set")
	return nil
}