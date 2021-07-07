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

	// error if tokens empty
	if len(tokens) == 0 {
		return nil, errors.New("Unexpected end of file")
	}

	head, tokens := tokens[0], tokens[1:]

	// First token cannot indicate the end of a list.
	if head == ")" {
		return nil, errors.New("Unexpected ')'")
	}

	// Token indicates the start of a list.
	if head == "(" {
		return ParseList(tokens), nil
	}

	// Token is either a Number or a Symbol
	return ParseAtom(head), nil
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

// ParseAtom first tries to parse a token as a Number, and if it fails, parses
// it as a Symbol.
func ParseAtom(s string) expression.Expression {
	maybeNum, err := ParseNumber(s)
	if err != nil {
		return ParseSymbol(s)
	}
	return maybeNum
}

// ParseList parses a list of tokens into a List.
func ParseList(tokens []string) expression.List {
	l := expression.NewList()
	for i := 0; i < len(tokens); {
		switch t := tokens[i]; t {
		case ")":
			return l
		case "(":
			sublist := ParseList(tokens[i+1:])
			l = l.Append(sublist)
			i += sublist.Length() + 2
		default:
			l = l.Append(ParseAtom(t))
			i++
		}
	}
	return l
}
