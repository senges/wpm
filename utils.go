/* Stolen from go-git repository : https://github.com/go-git/go-git/blob/master/_examples/common.go */
package main

import (
	"fmt"
	"os"
)

const VERBOSE = true

/* Get current working directory */
func getWD() string {
	path, err := os.Getwd()
	CheckIfError(err)

	return path
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	//os.Exit(1)
	panic(err)
}

func INFO(format string, args ...interface{}) {
	if VERBOSE {
		fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
	}
}

func CMD(format string, args ...interface{}) {
	if VERBOSE {
		fmt.Printf("\x1b[34;1m> %s\x1b[0m\n", fmt.Sprintf(format, args...))
	}
}

func OK(format string, args ...interface{}) {
	if VERBOSE {
		fmt.Printf("\x1b[36;1m%s\x1b[0m\n\n", fmt.Sprintf(format, args...))
	}
}