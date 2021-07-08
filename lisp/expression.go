package lisp

import ()

// An Expression is a List or an Atom. Lisp proceeds by evaluating Expressions.
type Expression interface {
	// Equal checks if an expression matches another Expression without
	// evaluating them.
	Equal(Expression) bool

	// String returns a readable Lisp string representing the Expression.
	String() string
}
