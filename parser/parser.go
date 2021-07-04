package parser

import (
	"strings"
)

// Parse tries to read a string into a Lisp expression.
func Parse(program string) (Expression, error) {

	return ReadFromTokens(Tokenize(program))
}

// Tokenize splits a string into tokens for the parser.
func Tokenize(program string) []string {

	// Pad parentheses with spaces, then split on spaces.
	return strings.Split(
		strings.ReplaceAll(
			strings.ReplaceAll(program, "(", " ( "),
			")", " ) ",
		),
		" ",
	)
}

// ReadFromTokens tries to parse a list of tokens into a Lisp expression.
func ReadFromTokens(tokens []string) (Expression, error) {
}
