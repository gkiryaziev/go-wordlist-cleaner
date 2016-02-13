package operations

import (
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb"
	"os"
	//"time"

	s "../service"
)

func isPrint(text string) bool {
	for _, r := range text {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}

func isSize(min, max int, line string) bool {
	if len([]rune(line)) < min || len([]rune(line)) > max {
		return false
	}
	return true
}

func DoClean(remove, trim bool, min, max int, src_file, new_file string) error {

	var added int64 = 0

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

	// Progress Bar
	bar := pb.New64(total)
	bar.ShowPercent = true
	bar.ShowBar = true
	bar.ShowCounters = true
	bar.ShowTimeLeft = true
	//bar.SetRefreshRate(time.Millisecond * 100)
	//bar.Format("<.- >")
	bar.Start()

	for scanner.Scan() {
		line := scanner.Text()

		if remove && trim {
			if isPrint(line) && isSize(min, max, line) {
				fmt.Fprintln(writer, line)
				added++
			}
		}

		if remove && !trim {
			if isPrint(line) {
				fmt.Fprintln(writer, line)
				added++
			}
		}

		if !remove && trim {
			if isSize(min, max, line) {
				fmt.Fprintln(writer, line)
				added++
			}
		}
		bar.Increment()
	}

	bar.Finish()

	if err := writer.Flush(); err != nil {
		return err
	}

	fmt.Println("\nResult:", src_file)
	fmt.Println("-------------------------------------------")
	fmt.Printf("|%-20s|%20d|\n", "Total", total)
	fmt.Printf("|%-20s|%20d|\n", "Removed", (total - added))
	fmt.Printf("|%-20s|%20d|\n", "Result", added)
	fmt.Println("-------------------------------------------\n")

	return scanner.Err()
}
