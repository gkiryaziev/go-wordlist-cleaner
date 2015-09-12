package operations

import (
	"bufio"
	"fmt"
	"os"

	s "../service"
)

func IsPrint(text string) bool {
	for _, r := range text {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}

func IsSize(min, max int, line string) bool {
	if len([]rune(line)) < min || len([]rune(line)) > max {
		return false
	}
	return true
}

func DoClean(remove, trim bool, min, max int, src_file, new_file string) error {

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

		if remove && trim {
			if IsPrint(line) && IsSize(min, max, line) {
				fmt.Fprintln(writer, line)
				added++
			}
		}

		if remove && !trim {
			if IsPrint(line) {
				fmt.Fprintln(writer, line)
				added++
			}
		}

		if !remove && trim {
			if IsSize(min, max, line) {
				fmt.Fprintln(writer, line)
				added++
			}
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
	fmt.Println("Cleaning result")
	fmt.Printf("|%-20s|%20d|\n", "Total", total)
	fmt.Printf("|%-20s|%20d|\n", "Removed", (total - added))
	fmt.Printf("|%-20s|%20d|\n", "Result", added)

	return scanner.Err()
}
