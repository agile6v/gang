package util

import (
	"fmt"
	"os"
)

const (
	EXIT_SUCCESS = iota
    EXIT_ERROR
	EXIT_BAD_ARGS
)

func ReturnError(code int, err error) {
	fmt.Fprintln(os.Stderr, "Error: ", err)
	os.Exit(code)
}

