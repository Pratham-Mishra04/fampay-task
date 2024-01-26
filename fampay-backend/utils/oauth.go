package utils

import (
	"encoding/gob"
	"log"
	"os"

	"golang.org/x/oauth2"
)

// Save token securely (e.g., in a database)
func SaveToken(token *oauth2.Token) {
	// Implement your token storage logic here
	// Example: Save the token to a file
	file, err := os.Create("token.json")
	if err != nil {
		log.Printf("Error saving token: %v", err)
		return
	}
	defer file.Close()

	gob.NewEncoder(file).Encode(token)
}

// Load token from storage
func LoadToken() (*oauth2.Token, error) {
	// Implement your token retrieval logic here
	// Example: Load the token from a file
	file, err := os.Open("token.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var token oauth2.Token
	err = gob.NewDecoder(file).Decode(&token)
	return &token, err
}
