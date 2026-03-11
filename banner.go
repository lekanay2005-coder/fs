package main

import (
	"errors"
	"os"
	"strings"
)

// GenerateASCII creates ASCII art for the text using the banner file
func GenerateASCII(text, banner string) (string, error) {

	// Open the banner file like "standard.txt"
	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		return "", errors.New("banner file not found")
	}

	// Split the file into lines
	lines := strings.Split(string(data), "\n")

	// This will store the final ASCII art
	result := ""

	// Each ASCII letter uses 8 rows
	for i := 0; i < 8; i++ {

		// Loop through each character in the text
		for _, char := range text {

			// Find the starting line of the character in the banner
			// ASCII printable characters start at code 32 (space)
			index := (int(char) - 32) * 9 // 8 lines + 1 empty line per char

			// Add this row of the character to result
			result += lines[index+i]
		}

		// Move to next line after finishing row
		result += "\n"
	}

	// Return the full ASCII art
	return result, nil
}