package main

import (
	"fmt"
	"math/rand"
)

func compare (random, guess int) bool {
	if guess < 1 {
		fmt.Println("Exiting game...")
		return true
	} else if guess == random {
		fmt.Println("You win!")
		return true
	} else if guess > random {
		fmt.Println("Too big!")
		return false
	} else {
		fmt.Println("Too small!")
		return false
	}
}
        
func main() {
    max := 10
	random := rand.Intn(max - 1) + 1

	var guess int

	for {
		fmt.Print("Guess the number from one to ten: ")
		fmt.Scanf("%d", &guess)
		if compare(random, guess) == true {
			break
		}
	}

	// for comparison := false; !comparison; {
	// 	fmt.Print("Guess the number from one to ten: ")
	// 	fmt.Scanf("%d", &guess)
	// 	comparison = compare(random, guess)
	// }
} 
