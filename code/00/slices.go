package main

import "fmt"

func main() {
	var x = [...]int{1, 2, 3, 4}
	var slice = x[:2]
	fmt.Println("x, slice:", x, slice)
	x[0] = 42
	fmt.Println("slice:", slice)
	fmt.Println("len, cap:", len(slice), cap(slice))
	slice = slice[:cap(x)]
	fmt.Println("slice:", slice)
}
