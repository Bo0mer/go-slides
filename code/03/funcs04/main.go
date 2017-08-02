package main

import (
	"errors"
	"fmt"
)

// START OMIT
func theAnswer() (answer int, err error) {
	if answer, err = readRemoteAnswer(); err != nil {
		if !isTemporary(err) {
			return
		}
		answer = -42
	}

	localAnswer, otherErr := readLocalAnswer()
	if err != nil {
		return
	}
	answer, err = localAnswer, otherErr
	return
}

func main() {
	fmt.Println(theAnswer())
}

// END OMIT

func isTemporary(err error) bool { return true }

func readRemoteAnswer() (int, error) {
	return 0, errors.New("")
}

func readLocalAnswer() (int, error) {
	return 42, nil
}
