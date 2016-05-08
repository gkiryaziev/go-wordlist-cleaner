package operations

import (
	"fmt"

	s "github.com/gkiryaziev/go-wordlist-cleaner/service"
)

// DoCalculate calculate lines in source file
func DoCalculate(srcFile string) error {

	total, err := s.CalculateLines(srcFile)
	if err != nil {
		return err
	}

	fmt.Printf("|%-40s|%20d|\n", srcFile, total)

	return nil
}
