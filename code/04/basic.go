package main

import (
	"log"

	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatal("failed to start Chrome:", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("failed to open page:", err)
	}

	page.Size(1024, 768)
	page.Navigate("https://google.com")
}
