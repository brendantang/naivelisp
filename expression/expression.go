package expression

import (
	"fmt"
	"strconv"
)

// An Expression is a List or an Atom. Lisp proceeds by evaluating expressions.
type Expression interface {
}

// A Symbol is an Atomic value which represents a variable name.
type Symbol struct {
	name string
}

// A Number is an Atomic value representing a number. All numbers are internally
// Go float64.
type Number struct {
	value float64
}

// NewNumber creates a Number from a string.
func NewNumber(s string) (Number, error) {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return Number{0}, fmt.Errorf("could not convert '%s' into a number", s)
	}
	return Number{val}, nil

}

// NumberFromFloat constructs a Number.
func NumberFromFloat(f float64) Number {
	return Number{f}
}

// NewSymbol creates a Symbol from a string.
func NewSymbol(s string) Symbol {
	return Symbol{s}
}

// A List is a list of expressions. A list is represented internally as a Go
// slice.
type List struct {
	elements []Expression
}

// NewList creates a List from a slice of Expressions.
func NewList(es ...Expression) List {
	return List{es}
}

// ToSlice converts a list into a slice of Expressions.
func (l List) ToSlice() []Expression {
	return l.elements
}

// An Environment maps variable names to their values.
type Environment interface{}

// Eval evaluates a Lisp expression in an environment.
func Eval(x Expression, env Environment) (Expression, Environment, error) {
	return nil, nil, nil
}

// Equal checks if two expressions are equal without evaluating
// them.
func Equal(a, b Expression) bool {
	return a == b
}
