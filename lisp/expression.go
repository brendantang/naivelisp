package lisp

import ()

// An expression is a List or an Atom. Lisp proceeds by evaluating expressions.
type expression interface {

	// String returns a readable Lisp string representing the expression.
	String() string
}
