package expression

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