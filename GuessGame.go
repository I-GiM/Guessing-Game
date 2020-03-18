package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Can Your Guess Save Your Life!?")
	println()

	//generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)

	var guess int
	/*fmt.Scan(&guess)

	fmt.Printf("You guessed %d\n", guess)

	if guess == secretNumber {
		fmt.Println("Subarashi! You live!")
	} else {
		fmt.Println("Omae wa mo shindeiru")
	}*/
	for i := 0; i < 3; i++ {
		fmt.Println("Input your guess")
		fmt.Scan(&guess)
		if guess == secretNumber {
			fmt.Println("Subarashi! You Live!!!")
			break
		}
		if guess != secretNumber {
			fmt.Println("Wrong! Try again")
			println()
		}
		if i == 2 {
			fmt.Println("Wrong!!! Omae wa mo shindeiru")
			println()
		}
	}
}
