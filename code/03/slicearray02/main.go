package main

import "fmt"

func main() {
	slice := make([]int, 10)

	for i := range slice {
		slice = append(slice, i) // HL
	}

	fmt.Println(len(slice))
}
