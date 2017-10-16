package testing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

const pageURL = "file:///Users/ivan/workspace/go-slides/code/05/testing/page.html"

// START1 OMIT
var _ = Describe("My Go web page", func() {
	var page *agouti.Page
	BeforeEach(func() {
		var err error
		page, err = driver.NewPage()
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
