package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello, enter your age: ")
	var age int
	fmt.Scanf("%d", &age)
	months := age * 12
	days := age * 365
	fmt.Println(months, "months, ", days, "days")
}
