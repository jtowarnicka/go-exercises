package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var secretNumber int;

func randomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func compare(guess string) bool {
	if guess == "exit" {
		fmt.Println("Exiting game...")
		return true
	} else {
		guessInt, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println("Invalid input")
			return false
		} else if guessInt == secretNumber {
			fmt.Println("You win!")
			fmt.Println("Do you want to play again? (y/n)")
			var answer string
			fmt.Scan(&answer)
			if answer == "y" {
				fmt.Println("Starting a new game...")
				secretNumber = randomNumber(1, 10)
				return false
			} else {
				fmt.Println("Exiting game...")
				return true
			}
		} else if guessInt > secretNumber {
			fmt.Println("Too big!")
			return false
		} else {
			fmt.Println("Too small!")
			return false
		}
	}
}

func main() {
	secretNumber = randomNumber(1, 10)
	var guess string
	fmt.Println(secretNumber)

	fmt.Println("Guess the number from one to ten. If you want to exit, enter \"exit\".")
	for check := false; !check; {
		fmt.Print("Enter the number: ")
		fmt.Scan(&guess);

		check = compare(guess)
	}
}