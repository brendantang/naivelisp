package lisp

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	for _, c := range parserTestCases {
		tokens := Tokenize(c.input)
		if !TokensEqual(tokens, c.expectedTokens) {
			t.Fatalf("FAIL: %s\nInput: %q\nExpected: %#v\nGot: %#v\n", c.description, c.input, c.expectedTokens, tokens)
		}
	}
}

func TestParse(t *testing.T) {
	for _, c := range parserTestCases {
		exp, err := Parse(c.input)
		if err != nil {
			t.Fatalf("FAIL: %s\nError: %v", c.description, err)
		}
		if exp.String() != c.expectedExpression.String() {
			t.Fatalf("FAIL: %s\nInput: %q\nExpected: %#v\nGot: %#v\n", c.description, c.input, c.expectedExpression.String(), exp.String())
		}
	}
}

// TokensEqual tells whether two lists of strings (tokens) are equal.
func TokensEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true

}

var parserTestCases = []struct {
	description        string
	input              string
	expectedTokens     []string
	expectedExpression Expression
}{
	{
		description:        "symbol",
		input:              "foo",
		expectedTokens:     []string{"foo"},
		expectedExpression: NewSymbol("foo"),
	},
	{
		description:        "integer",
		input:              "25",
		expectedTokens:     []string{"25"},
		expectedExpression: NewNumber(25),
	},
	{
		description:        "float",
		input:              "253254.3",
		expectedTokens:     []string{"253254.3"},
		expectedExpression: NewNumber(253254.3),
	},
	{
		description:    "list",
		input:          "(1 3 4 5)",
		expectedTokens: []string{"(", "1", "3", "4", "5", ")"},
		expectedExpression: NewList(
			NewNumber(1),
			NewNumber(3),
			NewNumber(4),
			NewNumber(5),
		),
	},
	{
		description:    "nested lists",
		input:          "(1 (2 3) (4 (5)))",
		expectedTokens: []string{"(", "1", "(", "2", "3", ")", "(", "4", "(", "5", ")", ")", ")"},
		expectedExpression: NewList(
			NewNumber(1),
			NewList(
				NewNumber(2),
				NewNumber(3),
			),
			NewList(
				NewNumber(4),
				NewList(
					NewNumber(5),
				),
			),
		),
	},
	{
		description:    "extra spaces",
		input:          "(1 (2 3  )   (4  (5)))",
		expectedTokens: []string{"(", "1", "(", "2", "3", ")", "(", "4", "(", "5", ")", ")", ")"},
		expectedExpression: NewList(
			NewNumber(1),
			NewList(
				NewNumber(2),
				NewNumber(3),
			),
			NewList(
				NewNumber(4),
				NewList(
					NewNumber(5),
				),
			),
		),
	},
}
