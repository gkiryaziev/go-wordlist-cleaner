package operations

import (
	"bufio"
	"fmt"
	"os"

	s "../service"
)

// Search duplicates in source file and write uniq to new file
func DoDuplicate(src_file, new_file string) error {

	m := map[uint64]bool{}
	counter := 0
	percent := 0
	added := 0
	total, err := s.CalculateLines(src_file)
	if err != nil {
		return err
	}

	in, err := os.Open(src_file)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(new_file)
	if err != nil {
		return err
	}
	defer out.Close()

	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	fmt.Printf("\n%s processing: ", src_file)

	for scanner.Scan() {
		line := scanner.Text()
		line_hash := s.GetHashFvn64(line)

		if _, seen := m[line_hash]; !seen {
			fmt.Fprintln(writer, line)
			m[line_hash] = true
			added++
		}

		counter++
		if counter == 100000 {
			percent += counter
			fmt.Printf("..%d%%", (percent * 100 / total))
			counter = 0
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Duplicate search result")
	fmt.Printf("|%-20s|%20d|\n", "Total", total)
	fmt.Printf("|%-20s|%20d|\n", "Removed", (total - added))
	fmt.Printf("|%-20s|%20d|\n", "Result", added)

	return scanner.Err()
}
