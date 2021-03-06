package lisp

import (
	"errors"
	"fmt"
)

// runtimeError describes an error encountered while evaluating Lisp code.
func runtimeError(description string) error {
	return errors.New(description)
}

// runtimeErrorf describes an error encountered while evaluating Lisp code using
// a format string.
func runtimeErrorf(formatString string, vals ...interface{}) error {
	return runtimeError(fmt.Sprintf(formatString, vals...))
}

// typeError describes a type error.
func typeError(val Expression, want string) error {
	return fmt.Errorf("type error: %T is not a %s", val, want)
}
