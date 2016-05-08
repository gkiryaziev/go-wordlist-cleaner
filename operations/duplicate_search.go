package operations

import (
	"bufio"
	"fmt"
	"os"
	//"time"

	"github.com/cheggaaa/pb"

	s "github.com/gkiryaziev/go-wordlist-cleaner/service"
)

// DoDuplicate search duplicates in source file and write uniq to new file
func DoDuplicate(srcFile, newFile string) error {

	m := map[uint64]bool{}

	var added int64

	total, err := s.CalculateLines(srcFile)
	if err != nil {
		return err
	}

	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(newFile)
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
		lineHash := s.GetHashFvn64(line)

		if _, seen := m[lineHash]; !seen {
			fmt.Fprintln(writer, line)
			m[lineHash] = true
			added++
		}
		bar.Increment()
	}

	bar.Finish()

	if err := writer.Flush(); err != nil {
		return err
	}

	fmt.Println("\nResult:", srcFile)
	fmt.Println("-------------------------------------------")
	fmt.Printf("|%-20s|%20d|\n", "Total", total)
	fmt.Printf("|%-20s|%20d|\n", "Removed", (total - added))
	fmt.Printf("|%-20s|%20d|\n", "Result", added)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	return scanner.Err()
}
