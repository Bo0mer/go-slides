package main

import "os"

func main() {
}

// START OMIT
func save(path string, data []byte) error {
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close() // HL

	_, err = fd.Write(data)
	return err
}

// END OMIT
