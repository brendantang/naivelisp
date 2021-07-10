package lisp

type procedure struct {
	name  string
	body  func(...expression) (expression, error)
	arity int
}

func (p procedure) call(xs ...expression) (expression, error) {
	if len(xs) != p.arity {
		return nil, runtimeErrorf(
			"tried to apply procedure '%s' to %d arguments, but '%s' takes %d arguments",
			p.name,
			len(xs),
			p.name,
			p.arity,
		)
	}
	return p.body(xs...)
}

func newProcedure(name string, arity int, body func(...expression) (expression, error)) procedure {
	return procedure{
		name:  name,
		arity: arity,
		body:  body,
	}
}

//String should represent proc as a string
func (p procedure) String() string {
	return "#procedure"
}
