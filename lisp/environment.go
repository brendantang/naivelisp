package lisp

// An Environment binds variable names to their values.
type Environment struct {
	bindings map[string]Expression
}

func (e Environment) get(key string) (exp Expression, ok bool) {
	exp, ok = e.bindings[key]
	return
}

// StdEnv returns an Environment which binds variable names for the standard,
// primitive values (like `+` and `-`).
func StdEnv() Environment {
	bindings := map[string]Expression{
		"+": newProcedure(
			"+",
			2,
			func(xs ...Expression) (Expression, error) {
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
			func(xs ...Expression) (Expression, error) {
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
	return Environment{bindings}
}
