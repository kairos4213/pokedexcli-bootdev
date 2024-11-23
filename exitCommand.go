package main

import (
	"os"
)

func exitCommand(c *config) error {
	os.Exit(0)
	return nil
}
