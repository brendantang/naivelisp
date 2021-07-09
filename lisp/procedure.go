package lisp

type procedure struct {
	body func(...Expression) Expression
}

// Equal should checck if two procedures are equivalent
func (p procedure) Equal(other Expression) bool {
	panic("TODO: implement comparing procedures")
}

//String should represent proc as a string
func (p procedure) String() string {
	return "some procedure"
}
