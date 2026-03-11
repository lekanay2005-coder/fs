package main

import "fmt"

// main function runs everything
func main() {

	// ParseArgs() reads what the user typed
	text, banner, err := ParseArgs()

	// If the arguments are wrong, show usage
	if err != nil {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return // stop the program
	}

	// GenerateASCII() builds the ASCII art using the banner
	ascii, err := GenerateASCII(text, banner)

	// If something went wrong (e.g., banner file missing)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the ASCII art to the terminal
	fmt.Print(ascii)
}