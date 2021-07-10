package lisp

import ()

// An Expression is a List or an Atom. Lisp proceeds by evaluating Expressions.
type Expression interface {

	// String returns a readable Lisp string representing the Expression.
	String() string
}
