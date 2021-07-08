package lisp

import (
	"errors"
	"fmt"
)

// Eval evaluates a Lisp expression in an environment.
func Eval(x Expression, env Environment) (Expression, Environment, error) {

	switch exp := x.(type) {

	case Symbol:
		// look up the symbol in env
		val, keyDefined := env[exp.name]

		// handle undefined
		if !keyDefined {
			return val, env, RuntimeError(fmt.Sprintf("Tried to evaluate undefined symbol '%v'", exp.name))
		}

		// evaluate the symbol to its corresponding value defined in env
		return val, env, nil

	case Number:
		// numbers evaluate to themselves
		return exp, env, nil

	case List:
		elems := exp.elements

		// handle the `if` and `define` primitive
		procName, ok := elems[0].(Symbol)

		if !ok {
			return exp, env, RuntimeError(fmt.Sprintf("Tried to call a procedure with a non-symbol name ('%v')", elems[0]))
		}

		switch procName.name {

		case "if":
			return nil, nil, RuntimeError("TODO: implement `if`")
		case "define":
			return nil, nil, RuntimeError("TODO: implement `define`")
		default:
			// handle a procedure call
			return nil, nil, RuntimeError("TODO: implement procedure calls")

		}

	}

	return nil, nil, nil
}

// RuntimeError describes an error encountered while evaluating Lisp code.
func RuntimeError(description string) error {
	return errors.New(description)
}
