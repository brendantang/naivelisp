package expression

import ()

// An Expression is a List or an Atom. Lisp proceeds by evaluating expressions.
type Expression interface {
	// Equal checks if an expression matches another expression without
	// evaluating them.
	Equal(Expression) bool
}

// An Environment maps variable names to their values.
type Environment struct {
	vars map[string]Expression
}

// Eval evaluates a Lisp expression in an environment.
func Eval(x Expression, env Environment) (Expression, Environment, error) {
	return nil, env, nil
}
