package main

import (
	"fmt"
)

func HandlerAddProduct(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	return nil
}
