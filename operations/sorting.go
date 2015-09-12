package operations

import (
	"fmt"
	"sort"

	s "../service"
)

// Read source file, sort it alphabetically and write to new file
func DoSorting(src_file, new_file string) error {

	total, err := s.CalculateLines(src_file)
	if err != nil {
		return err
	}

	// read file
	fmt.Println("\nReading", src_file)
	source, err := s.ReadLine(src_file)
	if err != nil {
		return err
	}

	// sorting
	fmt.Println("Sorting", new_file)
	sort.Strings(source)
	fmt.Println(new_file, "sorted.")

	// write file
	fmt.Println("Saving", new_file)
	err = s.WriteLine(source, new_file)
	if err != nil {
		return err
	}

	fmt.Println(new_file, "saved.")

	fmt.Println()
	fmt.Println("Sorting result")
	fmt.Printf("|%-20s|%20d|\n", "Total", total)

	return nil
}
