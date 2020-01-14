package main

import "fmt"

// START OMIT
func main() {
	var bytes []byte
	var interass interface{}

	fmt.Println(bytes == nil)
	interass = bytes
	fmt.Println(interass == nil)
}

// END OMIT
