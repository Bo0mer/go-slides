package main

import (
	"fmt"
	"time"
)

// START OMIT
type Item struct {
	ID   int
	Name string
}

func main() {
	items := make([]*Item, 10)

	for i := range items {
		items[i] = &Item{ID: i}
	}

	for _, item := range items {
		go func() {
			fmt.Println(item.ID)
		}()
	}

	time.Sleep(time.Millisecond)
}

// END OMIT
