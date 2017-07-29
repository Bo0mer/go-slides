package main

import "fmt"

func main() {
	var m map[string]bool

	fmt.Println(m["FOO"])

	// You can take the len of a nil map.
	fmt.Println(len(m))

	// You can range over a nil map.
	for k, v := range m {
		fmt.Println(k, v)
	}
}
