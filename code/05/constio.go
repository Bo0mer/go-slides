package main

// START OMIT
type Error string

func (e Error) Error() string { return string(e) }

const EOF = Error("EOF") // HL
// END OMIT
