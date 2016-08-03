package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 10
	i := <-c
	fmt.Println(i)

	close(c)
	i, ok := <-c
	fmt.Println(i, ok)
	c <- 11 // runtime panic, send on closed channel
}
