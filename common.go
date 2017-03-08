// This file contains common helper functions.
package main

import (
	"errors"
	"fmt"
	"os"
)

// GetFileContents will read contents of file and return a byte array.
func GetFileContents(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Error opening file")
	}

	info, err := f.Stat()
	if err != nil {
		return nil, errors.New("Error getting file stats")
	}

	len := info.Size()
	data := make([]byte, len)
	n, err := f.Read(data)
	if err != nil {
		return nil, errors.New("Error reading file")
	}
	if int64(n) != len {
		return nil, errors.New("Could not read entire contents of file")
	}

	return data, nil
}

// PutFileContents will write given contents to given file.
func PutFileContents(filename string, contents string) error {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return errors.New("Error determining file permissions for writing")
	}

	f, err := os.OpenFile(filename, os.O_WRONLY, fileInfo.Mode().Perm())
	if err != nil {
		return errors.New("Error opening file to write")
	}

	_, err = f.WriteString(contents)
	if err != nil {
		return errors.New("Error writing file")
	}

	return nil
}

// Verbose prints debug messages if verbose flag is enabled in command-line options.
func Verbose(format string, a ...interface{}) {
	if flags.verbose {
		fmt.Printf(format+"\n", a...)
	}
}

// FailOnError is a helper function to validate err object and exit if necessary.
func FailOnError(err error, msg string) {
	if err != nil {
		os.Stderr.WriteString(msg + " : " + err.Error() + "\n")
		os.Exit(1)
	}
}

// Die is a helper function to exit the program abnormally.
func Die(format string, a ...interface{}) {
	os.Stderr.WriteString(fmt.Sprintf(format, a...) + "\n")
	os.Exit(1)
}
