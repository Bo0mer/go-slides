package main

import "fmt"

func min(x, y int) int {
	a := make([]struct{}, x)
	b := make([]struct{}, y)
	return copy(a, b) // HL
}

func main() {
	fmt.Println(min(10, 5))
}
