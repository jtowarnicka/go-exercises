package main

import "fmt"

func main() {
	fmt.Println("Hello, enter factors of a quadratic equasion: ")
	var a, b, c float64
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Print("c: ")
	fmt.Scan(&c)

	if (a == 0) {
		fmt.Println("Not a quadratic equasion")
	} else {
		delta := b*b - 4*a*c
		if delta < 0 {
			fmt.Println("No real roots")
		}
		if delta == 0 {
			fmt.Println("x0 = ", -b/(2*a))
		}
		if delta > 0 {
			fmt.Println("x1 =", (-b+delta)/(2*a), "x2 =", (-b-delta)/(2*a))
		}
	}
}
