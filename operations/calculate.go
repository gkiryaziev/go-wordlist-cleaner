package operations

import (
	"fmt"

	s "github.com/gkiryaziev/go-wordlist-cleaner/service"
)

// DoCalculate calculate lines in source file
func DoCalculate(src_file string) error {

	total, err := s.CalculateLines(src_file)
	if err != nil {
		return err
	}

	fmt.Printf("|%-40s|%20d|\n", src_file, total)

	return nil
}
