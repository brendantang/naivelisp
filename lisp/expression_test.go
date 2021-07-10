package lisp

import (
	"testing"
)

func TestString(t *testing.T) {
	for _, c := range ExpressionTestCases {
		asString := c.Expression.String()
		if !(asString == c.expectedString) {
			t.Fatalf("FAIL: %s\nInput: %q\nExpected: %#v\nGot: %#v\n", c.description, c.Expression, c.expectedString, asString)
		}
	}
}

var ExpressionTestCases = []struct {
	description    string
	Expression     Expression
	expectedString string
}{
	{
		description:    "symbol",
		Expression:     newSymbol("foo"),
		expectedString: "foo",
	},
	{
		description:    "number",
		Expression:     newNumber(123),
		expectedString: "123",
	},
	{
		description:    "number with decimal",
		Expression:     newNumber(35256.66),
		expectedString: "35256.66",
	},
	{
		description: "list",
		Expression: newList(
			newNumber(1),
			newSymbol("ok"),
			newNumber(23.345),
		),
		expectedString: "(1 ok 23.345)",
	},
	{
		description: "nested list",
		Expression: newList(
			newNumber(1),
			newSymbol("ok"),
			newList(
				newNumber(5),
				newSymbol("foo"),
			),
			newList(
				newList(
					newSymbol("ok"),
					newSymbol("computer"),
				),
				newList(
					newNumber(1),
				),
			),
			newNumber(23.345),
		),
		expectedString: "(1 ok (5 foo) ((ok computer) (1)) 23.345)",
	},
}
