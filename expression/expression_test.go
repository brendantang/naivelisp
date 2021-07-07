package expression

import (
	"testing"
)

func TestString(t *testing.T) {
	for _, c := range testCases {
		asString := c.expression.String()
		if !(asString == c.expectedString) {
			t.Fatalf("FAIL: %s\nInput: %q\nExpected: %#v\nGot: %#v\n", c.description, c.expression, c.expectedString, asString)
		}
	}
}

func TestEqual(t *testing.T) {
}

var testCases = []struct {
	description    string
	expression     Expression
	expectedString string
}{
	{
		description:    "symbol",
		expression:     NewSymbol("foo"),
		expectedString: "foo",
	},
	{
		description:    "number",
		expression:     NewNumber(123),
		expectedString: "123",
	},
	{
		description:    "number with decimal",
		expression:     NewNumber(35256.66),
		expectedString: "35256.66",
	},
	{
		description: "list",
		expression: NewList(
			NewNumber(1),
			NewSymbol("ok"),
			NewNumber(23.345),
		),
		expectedString: "(1 ok 23.345)",
	},
	{
		description: "nested list",
		expression: NewList(
			NewNumber(1),
			NewSymbol("ok"),
			NewList(
				NewNumber(5),
				NewSymbol("foo"),
			),
			NewList(
				NewList(
					NewSymbol("ok"),
					NewSymbol("computer"),
				),
				NewList(
					NewNumber(1),
				),
			),
			NewNumber(23.345),
		),
		expectedString: "(1 ok (5 foo) ((ok computer) (1)) 23.345)",
	},
}
