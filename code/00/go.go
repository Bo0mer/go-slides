package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Hi there!")
	}()
}
