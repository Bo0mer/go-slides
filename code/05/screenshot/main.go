package main

import (
	"log"

	"github.com/sclevine/agouti"
)

const (
	firstImage = "uid_dimg_0"
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
	// START OMIT
	page.Navigate("https://google.com")

	q := page.FindByName("q")
	q.SendKeys("golang gopher image")
	q.Submit()

	page.Screenshot("screenshot.png")
	// END OMIT
}
