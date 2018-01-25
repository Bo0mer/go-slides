package main

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type Error string

func (e Error) Error() string { return string(e) }

const EOF = Error("EOF") // HL

func main() {
	// START OMIT
	fmt.Println(io.EOF == errors.New("EOF")) // From Go's io package.

	fmt.Println(EOF == Error("EOF"))
	// END OMIT
}
