package lisp

import (
	"strings"
)

// A List is an expression representing a list of expressions. A list is represented internally as a Go
// slice.
type List struct {
	elements []Expression
}

// NewList creates a List from Expressions.
func NewList(es ...Expression) List {
	return List{es}
}

// ToSlice converts a List into a slice of Expressions.
func (l List) ToSlice() []Expression {
	return l.elements
}

// Length returns the number of Expressions in the List.
func (l List) Length() int {
	return len(l.elements)
}

// Equal checks if a List l is equal to another Expression without evaluating
// them.
func (l List) Equal(other Expression) bool {
	otherList, ok := other.(List)
	if ok {
		if len(l.elements) != len(otherList.elements) {
			return false
		}
		for i, v := range l.elements {
			if !v.Equal(otherList.elements[i]) {
				return false
			}
		}
		return true
	}
	return false

}

//String returns a Lisp string representing the List.
func (l List) String() string {
	var strs []string
	for _, el := range l.elements {
		strs = append(strs, el.String())
	}

	return "(" + strings.Join(strs, " ") + ")"
}

// Concat adds List other to the end of List l.
func (l List) Concat(other List) List {
	es := append(l.elements, other.elements...)
	return NewList(es...)
}

// Append adds Expression e to the end of List l.
func (l List) Append(e Expression) List {
	es := append(l.elements, e)
	return NewList(es...)
}