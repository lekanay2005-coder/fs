package main

import (
	"errors"
	"os"
)

// ParseArgs reads the command line arguments
func ParseArgs() (string, string, error) {

	args := os.Args // get everything user typed

	// If user typed only 1 argument (just the text)
	if len(args) == 2 {
		text := args[1]        // the text to convert
		banner := "standard"   // default banner if none specified
		return text, banner, nil
	}

	// If user typed 2 arguments (text + banner)
	if len(args) == 3 {
		text := args[1]        // the text to convert
		banner := args[2]      // the banner name
		return text, banner, nil
	}

	// Wrong format (too few or too many arguments)
	return "", "", errors.New("bad args")
}