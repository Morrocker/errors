package errors

import (
	"errors"
	"fmt"
	"os"
)

var sep string = ">"

// New creates a new Error to be used in logging.
func New(path string, i interface{}) error {
	var txt string
	switch v := i.(type) {
	case error:
		txt = error(v).Error()
	case string:
		txt = v
	case nil:
		txt = "<nil> value given to create error. Fix issue."
	default:
		fmt.Printf("wrong type given to the errors.New() function > %v", i)
		txt = fmt.Sprintf("%v", i)
		os.Exit(1)
	}

	errMsg := fmt.Sprintf("%s: %s", path, txt)
	return errors.New(errMsg)
}

// Single creates a single un-asociated error for contextless use
func Single(msg string) error {
	return errors.New(msg)
}

// Extend takes in a previous error and passes it down prepending the current "function chain link"
func Extend(path string, e error) error {
	errMsg := fmt.Sprintf("%s > %s", path, e)
	return errors.New(errMsg)
}
