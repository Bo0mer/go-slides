package testing_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"

	"testing"
)

var driver *agouti.WebDriver

func TestTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing Suite")
}

var _ = BeforeSuite(func() {
	SetDefaultEventuallyTimeout(2 * time.Second)
	driver = agouti.ChromeDriver()
	Ω(driver.Start()).Should(Succeed())
})

var _ = AfterSuite(func() {
	Ω(driver.Stop()).Should(Succeed())
})
