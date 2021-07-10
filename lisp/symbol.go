package lisp

// A symbol is an Atomic value which represents a variable name.
type symbol struct {
	name string
}

// newSymbol creates a Symbol from a string.
func newSymbol(s string) symbol {
	return symbol{s}
}

//String returns a Lisp string representing the symbol.
func (s symbol) String() string {
	return s.name
}
