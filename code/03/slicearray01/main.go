package main

import "fmt"

type Item struct {
	ID   int
	Name string
}

func main() {
	items := make([]Item, 10)
	for i, item := range items {
		item.ID = i
		item.Name = fmt.Sprintf("item %d", i)
	}

	fmt.Printf("%+v\n", items)
}
