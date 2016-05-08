package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	o "github.com/gkiryaziev/go-wordlist-cleaner/operations"
	s "github.com/gkiryaziev/go-wordlist-cleaner/service"
)

// doJob is main function
func doJob(remove, trim, duplicate, sorting, calculate bool, min, max int, srcFile, newFile string) error {

	// Check operations
	if !remove && !trim && !duplicate && !sorting && !calculate {
		return errors.New("Not specified operations.")
	}

	// Cleaning
	if (remove || trim) && (!duplicate && !sorting && !calculate) {
		if err := o.DoClean(remove, trim, min, max, srcFile, newFile); err != nil {
			return err
		}
		return nil
	}

	//Duplicate search
	if duplicate && (!remove && !trim && !sorting && !calculate) {
		if err := o.DoDuplicate(srcFile, newFile); err != nil {
			return err
		}
		return nil
	}

	// Sorting
	if sorting && (!remove && !trim && !duplicate && !calculate) {
		if err := o.DoSorting(srcFile, newFile); err != nil {
			return err
		}
		return nil
	}

	// calculate
	if calculate && (!remove && !trim && !duplicate && !sorting) {
		if err := o.DoCalculate(srcFile); err != nil {
			return err
		}
		return nil
	}

	return errors.New("Incorrect combination of operations.")

}

// Split file by name and extension
func splitFileName(file string) (string, string) {
	ext := filepath.Ext(file)
	name := file[0 : len(file)-len(ext)]
	return name, ext
}

func main() {

	// variables
	remove := false
	trim := false
	duplicate := false
	sorting := false
	calculate := false
	min := 8
	max := 63
	srcFile := "Dict.dic"
	newFile := "Dict_cleaned.dic"
	auto := false
	fileExt := ".dic"
	version := "0.2.6"

	// args
	for k, arg := range os.Args {
		switch arg {
		case "-h":
			s.Usage()
			return
		case "-v":
			fmt.Println(version)
			return
		case "remove":
			remove = true
		case "trim":
			trim = true
		case "duplicate":
			duplicate = true
		case "sort":
			sorting = true
		case "calculate":
			calculate = true
		case "-min":
			err := s.CheckArgs(len(os.Args), k)
			s.CheckError(err)
			i, err := strconv.Atoi(os.Args[k+1])
			s.CheckError(err)
			min = i
		case "-max":
			err := s.CheckArgs(len(os.Args), k)
			s.CheckError(err)
			i, err := strconv.Atoi(os.Args[k+1])
			s.CheckError(err)
			max = i
		case "-src":
			err := s.CheckArgs(len(os.Args), k)
			s.CheckError(err)
			srcFile = os.Args[k+1]
		case "-new":
			err := s.CheckArgs(len(os.Args), k)
			s.CheckError(err)
			newFile = os.Args[k+1]
		case "-a":
			auto = true
		case "-ext":
			err := s.CheckArgs(len(os.Args), k)
			s.CheckError(err)
			fileExt = "." + os.Args[k+1]
		}
	}

	// start time
	start := time.Now()

	if auto {
		filesList, err := s.SearchFilesInDir(fileExt, "./")
		s.CheckError(err)
		fmt.Println()
		fmt.Println(len(filesList), "files found.")
		fmt.Println()
		for _, srcFile := range filesList {
			name, ext := splitFileName(srcFile)
			newFile := name + "_cleaned" + ext
			err = doJob(remove, trim, duplicate, sorting, calculate, min, max, srcFile, newFile)
			s.CheckError(err)
		}
	} else {
		fmt.Println()
		err := doJob(remove, trim, duplicate, sorting, calculate, min, max, srcFile, newFile)
		s.CheckError(err)
	}

	// elapsed time
	elapsed := time.Since(start)
	fmt.Println("\nElapsed time: ", elapsed)
}
