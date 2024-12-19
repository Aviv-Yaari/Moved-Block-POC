package provider

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomID() string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	// Generate a random ID as a number, then convert to a string
	randomID := rand.Intn(1e8)           // You can adjust the range based on your needs
	return fmt.Sprintf("%08d", randomID) // Return as a zero-padded string
}
