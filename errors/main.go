//
// Golang Workshop 2024
//

package main

import (
	"errors"
	"fmt"
	"os"
)

type CustomError struct {
	Code int
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Custom Error with code: %d", e.Code)
}

func processFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	return CustomError{Code: 42}
}

func main() {

	filename := "file.txt"
	err := processFile(filename)
	var customErr CustomError

	if errors.As(err, &customErr) {
		fmt.Printf("Custom Error: Code %d\n", customErr.Code)
	} else {
		fmt.Println("Not a custom error")
	}

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File not found")
	} else {
		fmt.Println("Not a file not found error")
	}
}

type CustomFileError struct {
	msg string
	err error
}

func (e *CustomFileError) Error() string {
	return fmt.Sprintf("%s: %v", e.msg, e.err)
}

func (e *CustomFileError) Unwrap() error {
	return e.err
}

func WrapErrNotExist(msg string) error {
	return &CustomFileError{msg, os.ErrNotExist}
}
