package lisp

// An Environment maps variable names to their values.
type Environment struct {
	bindings map[string]Expression
}

func (e Environment) get(key string) (exp Expression, ok bool) {
	exp, ok = e.bindings[key]
	return
}

func (e Environment) Equal(other Environment) bool {
	panic("TODO: check if Environments are equal")
}

func StdEnv() Environment {
	bindings := map[string]Expression{
		"+": newProcedure(
			"+",
			2,
			func(xs ...Expression) (Expression, error) {
				var result Number
				for _, n := range xs[:2] {
					n, ok := n.(Number)
					if !ok {
						return nil, typeError(n, "Number")
					}
					result = NewNumber(result.value + n.value)
				}
				return result, nil
			},
		),
		"-": newProcedure(
			"-",
			2,
			func(xs ...Expression) (Expression, error) {
				a, ok := xs[0].(Number)
				if !ok {
					return nil, typeError(a, "Number")
				}
				b, ok := xs[1].(Number)
				if !ok {
					return nil, typeError(b, "Number")
				}
				result := NewNumber(a.value - b.value)
				return result, nil
			},
		),
	}
	return Environment{bindings}
}
