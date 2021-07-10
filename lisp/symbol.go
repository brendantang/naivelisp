package lisp

// A Symbol is an Atomic value which represents a variable name.
type Symbol struct {
	name string
}

// NewSymbol creates a Symbol from a string.
func NewSymbol(s string) Symbol {
	return Symbol{s}
}

//String returns a Lisp string representing the symbol.
func (s Symbol) String() string {
	return s.name
}
