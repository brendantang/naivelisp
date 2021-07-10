package lisp

import (
	"strconv"
)

// A number is an Atomic value representing a number. All numbers are internally
// Go float64.
type number struct {
	value float64
}

// newNumber creates a number from a float64.
func newNumber(f float64) number {
	return number{f}
}

//String returns a Lisp string representing the number.
func (n number) String() string {
	return strconv.FormatFloat(n.value, 'f', -1, 64)
}
