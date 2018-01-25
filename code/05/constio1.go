package main

type Error string

func (e Error) Error() string { return string(e) }

const EOF = Error("EOF") // HL

func main() {
	// START OMIT
	EOF = Error("not EOF anymore")
	// END OMIT
}
