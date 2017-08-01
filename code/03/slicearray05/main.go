package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	for i, v := range &a { // HL
		a[3] = 100
		if i == 3 {
			fmt.Printf("a[%d] = %d\n", i, v)
		}
	}
}
