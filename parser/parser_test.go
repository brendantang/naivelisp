package parser

import (
	"github.com/brendantang/naivelisp/expression"
	"testing"
)

func TestTokenize(t *testing.T) {
	for _, c := range testCases {
		tokens := Tokenize(c.input)
		if !TokensEqual(tokens, c.expectedTokens) {
			t.Fatalf("FAIL: %s\nInput: %q\nExpected: %#v\nGot: %#v\n", c.description, c.input, c.expectedTokens, tokens)
		}
	}
}

func TestParse(t *testing.T) {
	for _, c := range testCases {
		exp, err := Parse(c.input)
		if err != nil {
			t.Fatalf("FAIL: %s\nError: %v", c.description, err)
		}
		if !exp.Equal(c.expectedExpression) {
			t.Fatalf("FAIL: %s\nInput: %q\nExpected: %#v\nGot: %#v\n", c.description, c.input, c.expectedExpression, exp)
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

var testCases = []struct {
	description        string
	input              string
	expectedTokens     []string
	expectedExpression expression.Expression
}{
	{
		description:        "symbol",
		input:              "foo",
		expectedTokens:     []string{"foo"},
		expectedExpression: expression.NewSymbol("foo"),
	},
	{
		description:        "integer",
		input:              "25",
		expectedTokens:     []string{"25"},
		expectedExpression: expression.NewNumber(25),
	},
	{
		description:        "float",
		input:              "253254.3",
		expectedTokens:     []string{"253254.3"},
		expectedExpression: expression.NewNumber(253254.3),
	},
	{
		description:    "list",
		input:          "(1 3 4 5)",
		expectedTokens: []string{"(", "1", "3", "4", "5", ")"},
		expectedExpression: expression.NewList(
			expression.NewNumber(1),
			expression.NewNumber(3),
			expression.NewNumber(4),
			expression.NewNumber(5),
		),
	},
	{
		description:    "nested lists",
		input:          "(1 (2 3) (4 (5)))",
		expectedTokens: []string{"(", "1", "(", "2", "3", ")", "(", "4", "(", "5", ")", ")", ")"},
		expectedExpression: expression.NewList(
			expression.NewNumber(1),
			expression.NewList(
				expression.NewNumber(2),
				expression.NewNumber(3),
			),
			expression.NewList(
				expression.NewNumber(4),
				expression.NewList(
					expression.NewNumber(5),
				),
			),
		),
	},
	{
		description:    "extra spaces",
		input:          "(1 (2 3  )   (4  (5)))",
		expectedTokens: []string{"(", "1", "(", "2", "3", ")", "(", "4", "(", "5", ")", ")", ")"},
		expectedExpression: expression.NewList(
			expression.NewNumber(1),
			expression.NewList(
				expression.NewNumber(2),
				expression.NewNumber(3),
			),
			expression.NewList(
				expression.NewNumber(4),
				expression.NewList(
					expression.NewNumber(5),
				),
			),
		),
	},
}
