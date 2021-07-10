package lisp

import (
	"errors"
	"fmt"
)

// Eval / apply structure heavily influenced by SICP chapter: https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book-Z-H-26.html#%_sec_4.1

// Eval evaluates a Lisp expression in an environment.
func Eval(x Expression, env Environment) (Expression, Environment, error) {

	switch exp := x.(type) {

	// a number evaluates to itself
	case Number:
		return exp, env, nil

	// a symbol is a variable name
	case Symbol:
		// look up the symbol in env
		val, keyDefined := env.get(exp.name)

		// handle undefined
		if !keyDefined {
			return val, env, RuntimeErrorf("Tried to evaluate undefined symbol '%v'", exp.name)
		}

		// evaluate the symbol to its corresponding value defined in env
		return val, env, nil

	//
	case List:
		elems := exp.elements

		sym, ok := elems[0].(Symbol)

		if !ok {
			return exp, env, RuntimeErrorf("Tried to call a procedure with a non-symbol name ('%v')", elems[0])
		}

		switch sym.name {

		// handle the primitive special forms
		case "if":
			return exp, env, RuntimeError("TODO: implement `if`")
		case "define":
			return exp, env, RuntimeError("TODO: implement `define`")
		default:
			// handle a procedure call
			val, ok := env.get(sym.name)
			if !ok {
				return exp, env, RuntimeErrorf("Tried to apply undefined procedure '%v'", sym.name)
			}
			proc, ok := val.(procedure)
			return apply(proc, env, elems[1:]...)
		}
	}

	return x, env, nil
}

func apply(proc procedure, env Environment, arguments ...Expression) (Expression, Environment, error) {
	return proc.body(arguments...), env, nil
}

// RuntimeError describes an error encountered while evaluating Lisp code.
func RuntimeError(description string) error {
	return errors.New(description)
}
func RuntimeErrorf(formatString string, vals ...interface{}) error {
	return RuntimeError(fmt.Sprintf(formatString, vals...))
}
