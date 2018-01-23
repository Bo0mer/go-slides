package main

import "os"

// START OMIT
func save2(path string, data []byte) (err error) {
	fd, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := fd.Close(); err == nil && cerr != nil {
			err = cerr // HL
		}
	}()

	_, err = fd.Write(data)
	return err
}

// END OMIT
