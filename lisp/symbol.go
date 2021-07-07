package lisp

// A Symbol is an Atomic value which represents a variable name.
type Symbol struct {
	name string
}

// NewSymbol creates a Symbol from a string.
func NewSymbol(s string) Symbol {
	return Symbol{s}
}

// Equal checks if a Symbol s is equal to another Expression without evaluating
// them.
func (s Symbol) Equal(other Expression) bool {
	otherSym, ok := other.(Symbol)
	if ok {
		return s == otherSym
	}
	return false
}

//String returns a Lisp string representing the symbol.
func (s Symbol) String() string {
	return s.name
}
