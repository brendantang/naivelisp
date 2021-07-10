package lisp

import (
	"strconv"
)

// A Number is an Atomic value representing a number. All numbers are internally
// Go float64.
type Number struct {
	value float64
}

// NewNumber creates a Number from a float64.
func NewNumber(f float64) Number {
	return Number{f}
}

//String returns a Lisp string representing the number.
func (n Number) String() string {
	return strconv.FormatFloat(n.value, 'f', -1, 64)
}
