package tests

import (
	"log"
	"os"

	"github.com/TJM/go-trello"
)

// Client is used in all tests
var Client *trello.Client
var err error

// Test initialization
func init() {
	key := os.Getenv("API_KEY")
	token := os.Getenv("API_TOKEN")
	if key == "" || token == "" {
		log.Fatal("ERROR: Environment vars API_KEY and API_TOKEN must be set for trello tests.")
	}
	Client, err = trello.NewAuthClient(key, &token)
	if err != nil {
		log.Fatal("Error setting up client.")
	}
}
