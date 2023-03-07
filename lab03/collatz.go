package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func collatz(n, counter int) int {
	if n == 1 {
		return counter
	} else if n % 2 == 0 {
		return collatz(n / 2, counter + 1)
	} else {
		return collatz(3 * n + 1, counter)
	}
}

func loadFromFile() {
	seqCounter := map[int]int{}

	file, err := os.Open("statistics.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			fmt.Println("Invalid line:", scanner.Text())
			continue
		}
		result, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Invalid results:", fields[1])
			continue
		}
		seqCounter[result]++
	}

	var max int;
	var maxKey int;
	for key, value := range seqCounter {
    	if value > max {
        	max = value
			maxKey = key
    	}
	}
	fmt.Println("Max occurences for", maxKey, ": ", max)
}

func main() {
	file, err := os.Create("statistics.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 1000; i <= 1000000; i+=10000 {
		result := collatz(i, 1)
		fmt.Fprintf(file, "%d %d\n", i, result)
	}

	loadFromFile()
}
