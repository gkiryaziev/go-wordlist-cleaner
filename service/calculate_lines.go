package service

import (
	"bytes"
	"io"
	"os"
)

// CalculateLines return count of lines
func CalculateLines(path string) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buf := make([]byte, 8192)
	var count int64 = 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return count, err
		}

		count += int64(bytes.Count(buf[:c], lineSep))

		if err == io.EOF {
			break
		}
	}
	return count, nil
}
