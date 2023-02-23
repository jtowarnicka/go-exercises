package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello, enter your age: ")
	var age int
	fmt.Scanf("%d", &age)

	fmt.Printf("Your age on Mercury: %.2f \n", float64(age)*0.2408467)
	fmt.Printf("Your age on Venus: %.2f \n", float64(age)*0.61519726)
	fmt.Println("Your age on Earth: ", age)
	fmt.Printf("Your age on Mars: %.2f \n", float64(age)*1.8808158)
	fmt.Printf("Your age on Jupiter: %.2f \n", float64(age)*11.862615)
	fmt.Printf("Your age on Saturn: %.2f \n", float64(age)*29.447498)
	fmt.Printf("Your age on Uranus: %.2f \n", float64(age)*84.016846)
	fmt.Printf("Your age on Neptune: %.2f \n", float64(age)*164.79132)	
}
