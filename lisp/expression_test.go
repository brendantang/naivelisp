package lisp

import (
	"testing"
)

func TestString(t *testing.T) {
	for _, c := range expressionTestCases {
		asString := c.expression.String()
		if !(asString == c.expectedString) {
			t.Fatalf("FAIL: %s\nInput: %q\nExpected: %#v\nGot: %#v\n", c.description, c.expression, c.expectedString, asString)
		}
	}
}

var expressionTestCases = []struct {
	description    string
	expression     expression
	expectedString string
}{
	{
		description:    "symbol",
		expression:     newSymbol("foo"),
		expectedString: "foo",
	},
	{
		description:    "number",
		expression:     newNumber(123),
		expectedString: "123",
	},
	{
		description:    "number with decimal",
		expression:     newNumber(35256.66),
		expectedString: "35256.66",
	},
	{
		description: "list",
		expression: newList(
			newNumber(1),
			newSymbol("ok"),
			newNumber(23.345),
		),
		expectedString: "(1 ok 23.345)",
	},
	{
		description: "nested list",
		expression: newList(
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
