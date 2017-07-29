package main

import "fmt"

type writeError int

func (*writeError) Error() string { return "oops" }

// START OMIT
func write() *writeError { // HL
	return nil
}

func writeAndLog() error {
	err := write()
	if err != nil {
		fmt.Printf("error during write: %v\n", err)
	}
	return err
}

func main() {
	err := writeAndLog()
	fmt.Println(err == nil)
}

// END OMIT
