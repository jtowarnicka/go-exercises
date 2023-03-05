package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var secretNumber, highscore int;
var tries, rounds int = 1, 1;
var name string;

type Score struct {
	name  string
	tries int
	rounds int
}
var scores []Score

func randomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func registerScore(tries, rounds int) {
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	flag := false
	for i, score := range scores {
		if score.name == name {
			if score.tries < tries {
				scores[i].tries = tries
				fmt.Println("New personal record!")
			}
			if score.tries < highscore {
				fmt.Println("New global record!")
			}
			scores[i].rounds++
			flag = true
			break
		}
	}

	if !flag {
		scores = append(scores, Score{name, tries, rounds})
	}
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

			registerScore(tries, rounds)

			fmt.Println("Do you want to play again? (y/n)")
			var answer string
			fmt.Scan(&answer)

			if answer == "y" {
				fmt.Println("Starting a new game...")
				secretNumber = randomNumber(1, 10)
				tries = 0
				return false
			} else {
				fmt.Println("Exiting game...")
				return true
			}

		} else if guessInt > secretNumber {
			fmt.Println("Too big!")
			tries++
			return false
		} else {
			fmt.Println("Too small!")
			tries++
			return false
		}
	}
}

func loadFromFile() {
	file, err := os.Open("HallOfFame.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	savedHighScore := false
    for scanner.Scan() {
        fields := strings.Fields(scanner.Text())
        if len(fields) != 3 {
            fmt.Println("Invalid line:", scanner.Text())
            continue
        }
        tries, err := strconv.Atoi(fields[1])
        if err != nil {
            fmt.Println("Invalid tries:", fields[1])
            continue
        }
        rounds, err := strconv.Atoi(fields[2])
        if err != nil {
            fmt.Println("Invalid rounds:", fields[2])
            continue
        }
        scores = append(scores, Score{fields[0], tries, rounds})

		if !savedHighScore {
			highscore = tries
			savedHighScore = true
		}
    }
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
        return
    }
}

func saveToFile(scores []Score) {
	file, err := os.Create("HallOfFame.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, score := range scores {
        fmt.Fprintf(file, "%s %d %d\n", score.name, score.tries, score.rounds)
    }
}

func scoreboard() {
	sort.Slice(scores, func(i, j int) bool {
        return scores[i].tries < scores[j].tries
    })
	fmt.Println("scoreboard:")
	for _, score := range scores {
		fmt.Println(score.name, score.tries, score.rounds)
	}

	saveToFile(scores)
}

func main() {
	loadFromFile()

	secretNumber = randomNumber(1, 10)
	var guess string

	fmt.Println("Guess the number from one to ten. If you want to exit, enter \"exit\".")
	for check := false; !check; {
		fmt.Print("Enter the number: ")
		fmt.Scan(&guess);

		check = compare(guess)
	}
	scoreboard()
}
