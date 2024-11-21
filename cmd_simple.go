package main

import (
	"os"
	"os/exec"
)

// commandClear clears the terminal screen
func commandClear() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	return nil
}

// commandExit exits the program
func commandExit() error {
	os.Exit(0)

	return nil
}
