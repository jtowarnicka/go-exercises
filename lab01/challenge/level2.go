package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func randomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func compare(guess string, secret int) bool {
	if guess == "exit" {
		fmt.Println("Exiting game...")
		return true
	} else {
		guessInt, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println("Invalid input")
			return false
		} else if guessInt == secret {
			fmt.Println("You win!")
			return true
		} else if guessInt > secret {
			fmt.Println("Too big!")
			return false
		} else {
			fmt.Println("Too small!")
			return false
		}
	}
}

func main() {
	secretNumber := randomNumber(1, 10)
	var guess string

	fmt.Println("Guess the number from one to ten. If you want to exit, enter \"exit\".")
	for check := false; !check; {
		fmt.Print("Enter the number: ")
		fmt.Scan(&guess);

		check = compare(guess, secretNumber)
	}
}