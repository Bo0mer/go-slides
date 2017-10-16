package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

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

	page, err := agouti.NewPage(driver.URL(), agouti.Desired(agouti.Capabilities{
		"chromeOptions": map[string][]string{
			"args": []string{
				"headless",
				// There is no GPU on our Ubuntu box.
				"disable-gpu",
				// Sandbox requires namespace permissions that we don't have
				// on a container.
				"no-sandbox",
			},
		},
	}))
	if err != nil {
		log.Fatal("failed to open page:", err)
	}

	page.Size(1024, 768)
	// START OMIT
	page.Navigate("https://google.com")

	q := page.FindByName("q")
	q.SendKeys("golang gopher image")
	q.Submit()

	img := page.FindByID(firstImage)
	printImage(img)
	// END OMIT
}

func printImage(s *agouti.Selection) {
	data, err := s.Attribute("src")
	if err != nil {
		log.Fatal(err)
	}

	headerBody := strings.Split(data, ",")
	if len(headerBody) != 2 {
		return
	}

	imgData, err := base64.StdEncoding.DecodeString(headerBody[1])
	if err != nil {
		return
	}
	tmpfile, err := ioutil.TempFile("", "img.png")
	if err != nil {
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(imgData)); err != nil {
		return
	}
	if err := tmpfile.Close(); err != nil {
		return
	}

	img2txt := exec.Command("img2txt", "-f", "ansi", tmpfile.Name())
	img2txt.Stdout = os.Stdout
	if err := img2txt.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error printing image: %v\n", err)
	}
}
