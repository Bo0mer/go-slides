package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Ivan"] = 10
	m["Others"] = 10000

	for key, value := range m {
		fmt.Println(key, value)
	}

	if georgi, ok := m["Georgi"]; ok {
		fmt.Println(georgi)
	}

	delete(m, "Ivan")
}
