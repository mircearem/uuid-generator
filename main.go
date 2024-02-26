package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// Define the character set
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	// Parse command line flags
	newFlag := flag.Bool("new", false, "Generate a new random string")
	flag.Parse()

	// If the new flag is provided, generate and print a random string
	if *newFlag {
		randomString := generateRandomString(20)
		fmt.Println(randomString)
	} else {
		fmt.Println("--new flag not provided")
	}
}

func generateRandomString(length int) string {
	// Seed the random number generator
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	// Create a byte slice of the required length
	randomBytes := make([]byte, length)

	// Populate the byte slice with random characters from the charset
	for i := 0; i < length; i++ {
		randomBytes[i] = charset[rng.Intn(len(charset))]
	}

	// Convert the byte slice to a string and return it
	return string(randomBytes)
}
