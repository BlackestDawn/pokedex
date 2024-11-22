package main

import (
	"os"
	"os/exec"
)

// commandClear clears the terminal screen
func commandClear(cfg *config) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}

// commandExit exits the program
func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}
