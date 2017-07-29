package main

// START OMIT
type Foo interface {
	Foo() string
}

type foo []byte

func (f foo) Foo() string { return "panic" }

func main() {
	m := make(map[Foo]bool)

	var concreteFoo Foo = foo{0x42}
	m[concreteFoo] = true
}

// END OMIT
