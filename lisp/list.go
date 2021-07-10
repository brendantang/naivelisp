package lisp

import (
	"strings"
)

// A list is an expression representing a list of expressions. A list is represented internally as a Go
// slice.
type list struct {
	elements []expression
}

// newList creates a list from expressions.
func newList(es ...expression) list {
	return list{es}
}

// ToSlice converts a list into a slice of expressions.
func (l list) toSlice() []expression {
	return l.elements
}

// Length returns the number of expressions in the list.
func (l list) length() int {
	return len(l.elements)
}

//String returns a Lisp string representing the list.
func (l list) String() string {
	var strs []string
	for _, el := range l.elements {
		strs = append(strs, el.String())
	}

	return "(" + strings.Join(strs, " ") + ")"
}

// concat adds list other to the end of list l.
func (l list) concat(other list) list {
	es := append(l.elements, other.elements...)
	return newList(es...)
}

// push adds expression e to the end of list l.
func (l list) push(e expression) list {
	es := append(l.elements, e)
	return newList(es...)
}
