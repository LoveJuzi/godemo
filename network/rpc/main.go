package main

import "fmt"

func main() {
	a := 2111.11
	b := 1576.24

	s := 0.0
	for i := 0; i < 158; i++ {
		s += a
		s += b - float64(10*(i+1))
	}

	fmt.Println(s)
}
