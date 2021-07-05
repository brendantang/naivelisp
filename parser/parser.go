package parser

import (
	"errors"
	"fmt"
	"github.com/brendantang/naivelisp/expression"
	"strconv"
	"strings"
)

// Parse tries to read a string into a Lisp expression.
func Parse(program string) (expression.Expression, error) { return ReadFromTokens(Tokenize(program)) }

// Tokenize splits a string into tokens for the parser.
func Tokenize(program string) []string {

	// Pad parentheses with spaces, then split on spaces.
	return strings.Fields(
		strings.ReplaceAll(
			strings.ReplaceAll(program, "(", " ( "),
			")", " ) ",
		),
	)
}

// ReadFromTokens tries to parse a list of tokens into a Lisp expression.
func ReadFromTokens(tokens []string) (expression.Expression, error) {
	if len(tokens) == 0 {
		return nil, errors.New("Unexpected end of file")
	}

	first, rest := tokens[0], tokens[1:]

	// First token cannot indicate the end of a list.
	if first == ")" {
		return nil, errors.New("Unexpected ')'")
	}

	// First token indicates the start of a list.
	if first == "(" {
		// Initialize an empty list
		var l []expression.Expression

		// Range over the remaining tokens
		for i, currentToken := range rest {

			// Current token indicates the end of the list.
			if currentToken == ")" {
				return expression.NewList(l...), nil
			}

			//
			expressionsToAppend, err := ReadFromTokens(rest[i:])
			if err != nil {
				return nil, err
			}
			l = append(l, expressionsToAppend)

		}
	}

	// Token is either a Number or a Symbol
	maybeNum, err := ParseNumber(first)

	// If not a Number, must be a Symbol
	if err != nil {
		return ParseSymbol(first), nil
	}

	return maybeNum, nil
}

// ParseNumber tries to parse a Number from a string.
func ParseNumber(s string) (expression.Number, error) {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return expression.NewNumber(0), fmt.Errorf("could not convert '%s' into a number", s)
	}
	return expression.NewNumber(val), nil
}

// ParseSymbol parses a string into a Symbol.
func ParseSymbol(s string) expression.Symbol {
	return expression.NewSymbol(s)
}
