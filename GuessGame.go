package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess the number")
	fmt.Println("Please input your guess")

	//generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)

	var guess int
	fmt.Scan(&guess)

	fmt.Printf("You guessed %d\n", guess)
}
