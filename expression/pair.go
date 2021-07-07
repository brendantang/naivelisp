package expression

/*
// A Pair is an Expression constructed from two other Expressions.
type Pair struct {
	first Expression
	rest  Expression
}

// Null specifically signals the end of a linked list.
type Null struct{}

// Equal checks if another Expression is also Null.
func (n Null) Equal(other Expression) bool {
	_, ok := other.(Null)
	return ok
}

// Cons constructs a new pair from two Expressions.
func Cons(first, rest Expression) Pair {
	return Pair{first, rest}
}

// Car returns the first element of the Pair.
func (p Pair) Car() Expression {
	return p.first
}

// Cdr returns the second element of the Pair.
func (p Pair) Cdr() Expression {
	return p.rest
}

// Equal checks if another Expression is equal to Pair p.
func (p Pair) Equal(other Expression) bool {
	otherPair, ok := other.(Pair)
	if ok {
		return p.Car().Equal(otherPair.Car()) && p.Cdr().Equal(otherPair.Cdr())
	}
	return false
}

// ListLength treats a Pair as the beginning of a linked list, and returns the
// length of that list.
func (p Pair) ListLength() int {
	switch cdr := p.Cdr().(type) {
	case Null:
		return 1
	case Pair:
		return 1 + cdr.ListLength()
	default:
		return 2
	}
}
*/
