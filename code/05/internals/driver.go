package main

import (
	"net/http"
	"time"
)

// START OMIT
type WebDriver struct {
	Timeout    time.Duration
	Debug      bool
	HTTPClient *http.Client
	service    driverService // HL
	sessions   []*Session
}

type driverService interface { // HL
	URL() string
	Start(debug bool) error
	Stop() error
	WaitForBoot(timeout time.Duration) error
}

// END OMIT

type Session struct{}
