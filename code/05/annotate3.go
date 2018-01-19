package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	err := startApp()
	// START OMIT
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error caused by:\n\t %v", errors.Cause(err)) // HL
		os.Exit(1)
	}
	// END OMIT
}

func startApp() error {
	_, err := configFromFile("config.yml")
	if err != nil {
		return errors.Wrap(err, "error reading config") // HL
	}
	return nil
}

func configFromFile(path string) (*config, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	var cfg config
	if err = json.NewDecoder(fd).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

type config struct {
	Foo string `json:"foo"`
}
