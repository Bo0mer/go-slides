package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

func (p *Person) Greet() {
	fmt.Println("Hello, I'm", p.Name)
}

type Greeter interface {
	Greet()
}

func greet(g Greeter) {
	g.Greet()
}

func main() {
	greet(&Person{21, "Ivan"})
}
