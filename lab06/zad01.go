package main

import (
	"flag"
	"fmt"
	"math"
)

type Coefficients struct {
	a, b, c float64
}
type Tuple struct {
	x, y interface{}
}

func calculate(s *Coefficients) (Tuple, bool) {
	if s.a != 0 {
		delta := s.b*s.b - 4*s.a*s.c
		if delta >= 0 {
			return Tuple{(-s.b+math.Sqrt(delta))/(2*s.a), (-s.b-math.Sqrt(delta))/(2*s.a)}, true
		}
	}
	return Tuple{0, 0}, false
}

func main() {
	a := flag.Float64("a", 0, "a number")
    b := flag.Float64("b", 0, "b number")
	c := flag.Float64("c", 0, "c number")
    flag.Parse()
	structure := Coefficients{*a, *b, *c}
	fmt.Println(calculate(&structure))
}
