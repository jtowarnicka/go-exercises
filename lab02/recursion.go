package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

func helper_factorial(n int, acc *big.Int) *big.Int {
	if n == 0 {
		return acc
	}
	return helper_factorial(n-1, acc.Mul(acc, big.NewInt(int64(n))))
}
func factorial(n int) *big.Int {
	return helper_factorial(n, big.NewInt(1))
}

func fibonacci(n int, callCounter map[int]int) int {
    callCounter[n]++
    if n < 2 {
        return n
    }
    return fibonacci(n-1, callCounter) + fibonacci(n-2, callCounter)
}

func transform(word string) string {
	characters := map[string]string{"ą": "a", "ć": "c", "ę": "e", "ł": "l", "ń": "n", "ó": "o", "ś": "s", "ź": "z", "ż": "z"}

    newWord := word
	newWord = strings.ToLower(newWord)
    for key, value := range characters {
        newWord = strings.ReplaceAll(newWord, key, value)
    }
	return newWord
}

func toAscii(str1, str2 string) []*big.Int {
	var asciiList []*big.Int

	str1 = transform(str1)
	str2 = transform(str2)

	word := str1[:3] + str2[:3]
	fmt.Println("nickname:", word)
    for _, c := range word {
        asciiList = append(asciiList, big.NewInt(int64(c)))
    }
	return asciiList
}

func contains(slice []*big.Int, number *big.Int) bool {
    str := number.String()
    for _, value := range slice {
        substring := value.String()
        if !strings.Contains(str, substring) {
            return false
        }
    }
    return true
}

func findClosestNumber(x int, callCounter map[int]int) int {
	weakNumber := 0
	difference := 514229
	for key, value := range callCounter {
		diff := int(math.Abs(float64(value - x)))
		if diff < difference {
			weakNumber = key
			difference = diff
		}
		if diff > difference {
			break
		}
	}
	return weakNumber
}

func main() {
	var name, surname string;
	fmt.Print("Enter your first name: ");
	fmt.Scan(&name);
	fmt.Print("Enter your last name: ");
	fmt.Scan(&surname);

	nickname := toAscii(name, surname)

	fmt.Println("Your nickname in ASCII is:", nickname)

	var n int = 1;
	for check := false; !check; {
		n++;
		check = contains(nickname, factorial(n))
	}
	fmt.Println("Your strong number is", n)

	callCounter := map[int]int{}

	fibonacci(30, callCounter)
	weakNumber := findClosestNumber(n, callCounter)
	fmt.Println("Your weak number is", weakNumber)

}
