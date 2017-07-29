package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("google.com")
	if err != nil {
		fmt.Println("Where's my Internet?")
		return
	}
	defer resp.Body.Close() // HL
}
