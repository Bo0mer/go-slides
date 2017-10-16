package testing_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var pageURL = "file:///Users/ivan/workspace/go-slides/code/05/testing/page.html"

func init() {
	if p := os.Getenv("STATIC_FILE_PATH"); p != "" {
		pageURL = "file://" + p
	}
}

// START1 OMIT
func testForBrowser(browserName string, newPage func() (*agouti.Page, error)) {
	Describe("My Go web page on"+browserName, func() {
		var page *agouti.Page
		BeforeEach(func() {
			var err error
			page, err = newPage()
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("when viewing it", func() {
			BeforeEach(func() {
				Ω(page.Navigate(pageURL)).Should(Succeed())
			})

			It("should display proper title", func() {
				t, err := page.Title()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(t).Should(Equal("Hello, Gophers"))
			})
			// END1 OMIT

			// START2 OMIT
			It("should display loading message", func() {
				loadingDIV := page.FindByID("loading")
				Ω(loadingDIV).Should(BeVisible())
				Ω(loadingDIV).Should(HaveText("Loading..."))
			})
			// END2 OMIT

			// START3 OMIT
			It("should display Gopher greeting", func() {
				greetingDIV := page.FindByID("greeting")
				Ω(greetingDIV).Should(BeVisible())
				Ω(greetingDIV).Should(HaveText("Hello, Gophers!"))
			})

			// END3 OMIT
		})
	})
}

var _ = Describe("Chrome", func() {
	f := func() (*agouti.Page, error) {
		return agouti.NewPage(driver.URL(), agouti.Desired(agouti.Capabilities{
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
	}

	testForBrowser("Chrome", f)
})
