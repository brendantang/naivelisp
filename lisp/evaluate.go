package lisp

import ()

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
			return val, env, runtimeErrorf("Tried to evaluate undefined symbol '%v'", exp.name)
		}

		// evaluate the symbol to its corresponding value defined in env
		return val, env, nil

	//
	case List:
		elems := exp.elements

		sym, ok := elems[0].(Symbol)

		if !ok {
			return exp, env, runtimeErrorf("Tried to call a procedure with a non-symbol name ('%v')", elems[0])
		}

		switch sym.name {

		// handle the primitive special forms
		case "if":
			return exp, env, runtimeError("TODO: implement `if`")
		case "define":
			return exp, env, runtimeError("TODO: implement `define`")
		default:
			// handle a procedure call
			val, ok := env.get(sym.name)
			if !ok {
				return exp, env, runtimeErrorf("Tried to apply undefined procedure '%v'", sym.name)
			}
			proc, ok := val.(procedure)
			if !ok {
				return exp, env, runtimeErrorf("Tried to apply non-procedure type '%T'", proc)
			}

			// evaluate operands first
			args := elems[1:]
			operands := make([]Expression, len(args))
			for i, arg := range args {
				var err error
				operands[i], env, err = Eval(arg, env)
				if err != nil {
					return exp, env, err
				}
			}
			return apply(proc, env, operands...)
		}
	}

	return x, env, nil
}

func apply(proc procedure, env Environment, arguments ...Expression) (Expression, Environment, error) {
	exp, err := proc.call(arguments...)
	return exp, env, err
}
