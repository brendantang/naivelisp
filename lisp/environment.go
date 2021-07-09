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
	return false
}

func StdEnv() Environment {
	return Environment{make(map[string]Expression)}
}
