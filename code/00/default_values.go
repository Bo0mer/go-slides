package main

import "fmt"

func main() {
	var i int
	var f float32
	var b bool
	var s string
	fmt.Printf("%d %f %t %q\n", i, f, b, s)

	var ptr *int
	fmt.Printf("%v\n", ptr)
}
