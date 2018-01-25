package main

import (
	"io"
)

// START OMIT
func main() {
	io.ErrUnexpectedEOF = nil // HL
	// Now imagine working with the encoding/json package.
}

// END OMIT
