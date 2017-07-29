package main

import (
	"errors"
	"fmt"
)

type something struct{}

func initSomething(s *something) error { return errors.New("oops") }

func newSomething() (s *something, err error) {
	if err := initSomething(s); err != nil {
		fmt.Printf("error during something init: %v\n", err)
	}
	return
}

func main() {
	_, err := newSomething()
	fmt.Println(err)
}
