package lisp

// An environment binds variable names to their values.
type environment struct {
	bindings map[string]expression
}

func (e environment) get(key string) (exp expression, ok bool) {
	exp, ok = e.bindings[key]
	return
}

func stdEnv() environment {
	bindings := map[string]expression{
		"+": newProcedure(
			"+",
			2,
			func(xs ...expression) (expression, error) {
				var result number
				for _, n := range xs[:2] {
					n, ok := n.(number)
					if !ok {
						return nil, typeError(n, "number")
					}
					result = newNumber(result.value + n.value)
				}
				return result, nil
			},
		),
		"-": newProcedure(
			"-",
			2,
			func(xs ...expression) (expression, error) {
				a, ok := xs[0].(number)
				if !ok {
					return nil, typeError(a, "Number")
				}
				b, ok := xs[1].(number)
				if !ok {
					return nil, typeError(b, "Number")
				}
				result := newNumber(a.value - b.value)
				return result, nil
			},
		),
	}
	return environment{bindings}
}
