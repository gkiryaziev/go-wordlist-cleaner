package service

import (
	"fmt"
	"os"
	"path/filepath"
)

// Usage menu
func Usage() {
	a := filepath.Base(os.Args[0])
	fmt.Println()
	fmt.Println("Usage:", a, "[MODE] [OPERATIONS] [OPTIONS]")
	fmt.Println()
	fmt.Println("    WordList cleaner.")
	fmt.Println()
	fmt.Println("    Remove non-printable words, trim words length, search duplicates,")
	fmt.Println("    sorting, words counting.")
	fmt.Println()
	fmt.Println("Operations:")
	fmt.Println("    remove           Remove non-printable words.")
	fmt.Println("    trim             Trim file by size.")
	fmt.Println("    duplicate        Search duplicates.")
	fmt.Println("    sort             Sorting.")
	fmt.Println("    calculate        Calculate lines in the specified file.")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("    -min  INT        Minimal word length. [8]")
	fmt.Println("    -max  INT        Maximum word length. [63]")
	fmt.Println("    -src  STR        Source wordlist file. [Dict.dic]")
	fmt.Println("    -new  STR        New wordlist file. [Dict_cleaned.dic]")
	fmt.Println()
	fmt.Println("Mode:")
	fmt.Println("    -a               Automatic processing of all files in a directory. [false]")
	fmt.Println("    -ext  STR        File extension. Only for automatic mode. [dic]")
	fmt.Println()
	fmt.Println("    -h               This help.")
	fmt.Println("    -v               Print version.")
}
