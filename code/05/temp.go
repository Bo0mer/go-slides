package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

func main() {
	// START OMIT
	type temporary interface {
		Temporary() bool
	}

	if err := ping("google.com"); err != nil {
		if v, ok := err.(temporary); ok && v.Temporary() { // HL
			fmt.Println("google.com is temporary unavailable")
			return
		}
		fmt.Println("google.com is unavailable forever :o")
	}
	// END OMIT
}

func ping(url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	_, err := ctxhttp.Get(ctx, http.DefaultClient, "http://google.com")
	return err
}
