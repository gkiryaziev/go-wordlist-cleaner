package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	o "./operations"
	s "./service"
)

// Do job
func DoJob(remove, trim, duplicate, sorting, calculate bool, min, max int, src_file, new_file string) error {

	// Check operations
	if !remove && !trim && !duplicate && !sorting && !calculate {
		return errors.New("Not specified operations.")
	}

	// Cleaning
	if (remove || trim) && (!duplicate && !sorting && !calculate) {
		if err := o.DoClean(remove, trim, min, max, src_file, new_file); err != nil {
			return err
		}
		return nil
	}

	//Duplicate search
	if duplicate && (!remove && !trim && !sorting && !calculate) {
		if err := o.DoDuplicate(src_file, new_file); err != nil {
			return err
		}
		return nil
	}

	// Sorting
	if sorting && (!remove && !trim && !duplicate && !calculate) {
		if err := o.DoSorting(src_file, new_file); err != nil {
			return err
		}
		return nil
	}

	// calculate
	if calculate && (!remove && !trim && !duplicate && !sorting) {
		if err := o.DoCalculate(src_file); err != nil {
			return err
		}
		return nil
	}

	return errors.New("Incorrect combination of operations.")

}

// Split file by name and extension
func SplitFileName(file string) (string, string) {
	ext := filepath.Ext(file)
	name := file[0 : len(file)-len(ext)]
	return name, ext
}

func main() {

	// variables
	var remove bool = false
	var trim bool = false
	var duplicate bool = false
	var sorting bool = false
	var calculate bool = false
	var min int = 8
	var max int = 63
	var src_file string = "Dict.dic"
	var new_file string = "Dict_cleaned.dic"
	var auto bool = false
	var file_ext = ".dic"
	var version = "0.2.4"

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
			src_file = os.Args[k+1]
		case "-new":
			err := s.CheckArgs(len(os.Args), k)
			s.CheckError(err)
			new_file = os.Args[k+1]
		case "-a":
			auto = true
		case "-ext":
			err := s.CheckArgs(len(os.Args), k)
			s.CheckError(err)
			file_ext = "." + os.Args[k+1]
		}
	}

	// start time
	start := time.Now()

	if auto {
		files_list, err := s.SearchFilesInDir(file_ext, "./")
		s.CheckError(err)
		fmt.Println()
		fmt.Println(len(files_list), "files found.")
		fmt.Println()
		for _, src_file := range files_list {
			name, ext := SplitFileName(src_file)
			new_file := name + "_cleaned" + ext
			err = DoJob(remove, trim, duplicate, sorting, calculate, min, max, src_file, new_file)
			s.CheckError(err)
		}
	} else {
		err := DoJob(remove, trim, duplicate, sorting, calculate, min, max, src_file, new_file)
		s.CheckError(err)
	}

	// elapsed time
	elapsed := time.Since(start)
	fmt.Println("\nElapsed time: ", elapsed)
}
