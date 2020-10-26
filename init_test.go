package trello

import (
	"log"
	"math/rand"
	"os"
	"time"
)

// client is used in all tests
var client *Client
var err error

// Test initialization
func init() {
	key := os.Getenv("API_KEY")
	token := os.Getenv("API_TOKEN")
	if key == "" || token == "" {
		log.Fatal("ERROR: Environment vars API_KEY and API_TOKEN must be set for trello tests.")
	}
	rand.Seed(time.Now().UnixNano())
	client, err = NewAuthClient(key, &token)
	if err != nil {
		log.Fatal("Error setting up client.")
	}
}
