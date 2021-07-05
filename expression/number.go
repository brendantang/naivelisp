package expression

// A Number is an Atomic value representing a number. All numbers are internally
// Go float64.
type Number struct {
	value float64
}

// NewNumber creates a Number from a float64.
func NewNumber(f float64) Number {
	return Number{f}
}

// Equal checks if a Number n is equal to another Expression without evaluating
// them.
func (n Number) Equal(other Expression) bool {
	otherNum, ok := other.(Number)
	if ok {
		return n == otherNum
	}
	return false
}
