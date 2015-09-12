package service

import (
	"errors"
	"fmt"
	"os"
)

// Main check error
func CheckError(err error) {
	if err != nil {
		Usage()
		fmt.Println()
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

// Check is file or directory exist
func CheckFile(file string) error {
	f, err := os.Stat(file)
	if err == nil {
		if f.IsDir() {
			return errors.New("Directory with name " + file + " is exists.")
		} else {
			return errors.New("File with name " + file + " is exists.")
		}
	}
	return nil
}

// Check key value
func CheckArgs(args_length, arg_index int) error {
	if args_length == (arg_index + 1) {
		return errors.New("Not specified key value.")
	}
	return nil
}
