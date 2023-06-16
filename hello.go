package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate a random number between 1 and 100
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	// Prompt the user for input
	fmt.Println("Guess the number between 1 and 100:")

	var guess int
	for {
		// Read user input
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		// Compare the guess with the target number
		if guess > target {
			fmt.Println("Too high. Try again!")
		} else if guess < target {
			fmt.Println("Too low. Try again!")
		} else {
			fmt.Println("Congratulations! You guessed it correctly!")
			break
		}
	}
}
