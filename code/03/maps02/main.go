package main

import "fmt"

func main() {
	m := map[string]int{
		"Ivan": 42,
	}

	age := &m["Ivan"]
	fmt.Println(age)
}
