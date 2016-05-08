package service

import (
	"errors"
	"fmt"
	"os"
)

// CheckError check error
func CheckError(err error) {
	if err != nil {
		Usage()
		fmt.Println()
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

// CheckFile check is file or directory exist
func CheckFile(file string) error {
	f, err := os.Stat(file)
	if err == nil {
		if f.IsDir() {
			return errors.New("Directory with name " + file + " is exists.")
		}
		return errors.New("File with name " + file + " is exists.")
	}
	return nil
}

// CheckArgs check key value
func CheckArgs(argsLength, argIndex int) error {
	if argsLength == (argIndex + 1) {
		return errors.New("Not specified key value.")
	}
	return nil
}
