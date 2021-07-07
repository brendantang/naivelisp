package expression

import ()

// An Expression is a List or an Atom. Lisp proceeds by evaluating Expressions.
type Expression interface {
	// Equal checks if an expression matches another Expression without
	// evaluating them.
	Equal(Expression) bool

	// String returns a readable Lisp string representing the Expression.
	String() string
}

// An Environment maps variable names to their values.
type Environment struct {
	vars map[string]Expression
}

// Eval evaluates a Lisp expression in an environment.
func Eval(x Expression, env Environment) (Expression, Environment, error) {
	return nil, env, nil
}
