package interpreter

// An Expression is a List or an Atom. Lisp proceeds by evaluating expressions.
type Expression interface{}

// An Environment maps variable names to their values.
type Environment interface{}

// Eval evaluates a Lisp expression in an environment.
func Eval(x Expression, env Environment) (Expression, error) {
}
