package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("google.com")
	defer resp.Body.Close() // HL
	if err != nil {
		fmt.Println("Where's my Internet?")
		return
	}
}
