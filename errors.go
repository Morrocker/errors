package errors

import (
	"errors"
	"fmt"
	"os"
)

var sep string = ">"

// Error is a type that implements the error interface and stores suberrors entered info.
type Error struct {
	msg         error
	path        string
	function    string
	breadcrumbs []string
}

var _ error = (*Error)(nil)

func (e Error) Error() (err string) {
	if len(e.breadcrumbs) == 1 {
		err = e.breadcrumbs[0]
	} else {
		for i, crumb := range e.breadcrumbs {
			if i == 0 {
				err = crumb
			} else {
				err = fmt.Sprintf("%s > %s", crumb, err)
			}
		}
	}
	err = fmt.Sprintf("%s:%s", err, e.msg)
	return
}

// New creates a new Error to be used in logging.
func (e *Error) New(i interface{}) {
	switch v := i.(type) {
	case error:
		e.msg = error(v)
	case string:
		e.msg = errors.New(v)
	default:
		e.msg = errors.New("")
		fmt.Println("wrong type given to the errors.New() function")
		os.Exit(1)
	}
	e.Extend()
	return
}

// Extend extends the error breadcrumb list
func (e *Error) Extend() {
	crumb := fmt.Sprintf("%s.%s", e.path, e.function)
	e.breadcrumbs = append(e.breadcrumbs, crumb)
}

// SetPath sets the filepath of the function that generated the error.
func (e *Error) SetPath(p string) {
	e.path = p
}

// SetFunc sets the function that generated the error.
func (e *Error) SetFunc(f string) {
	e.function = f
}
