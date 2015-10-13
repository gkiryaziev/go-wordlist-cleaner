package operations

import (
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb"
	"os"
	//"time"

	s "../service"
)

// Search duplicates in source file and write uniq to new file
func DoDuplicate(src_file, new_file string) error {

	m := map[uint64]bool{}

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
		line_hash := s.GetHashFvn64(line)

		if _, seen := m[line_hash]; !seen {
			fmt.Fprintln(writer, line)
			m[line_hash] = true
			added++
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
