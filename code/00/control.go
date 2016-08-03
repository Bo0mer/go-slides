package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			switch i % 3 {
			case 1:
				fmt.Println("4")
			default:
				fmt.Println(i)
			}
		}
	}
}
