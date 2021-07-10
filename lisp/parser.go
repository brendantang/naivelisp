package lisp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Parse tries to read a string into a Lisp
func Parse(program string) (Expression, error) { return readFromTokens(tokenize(program)) }

// tokenize splits a string into tokens for the parser.
func tokenize(program string) []string {

	// Pad parentheses with spaces, then split on spaces.
	return strings.Fields(
		strings.ReplaceAll(
			strings.ReplaceAll(program, "(", " ( "),
			")", " ) ",
		),
	)
}

// readFromTokens tries to parse a list of tokens into a Lisp
func readFromTokens(tokens []string) (Expression, error) {

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
		return parseList(tokens), nil
	}

	// Token is either a number or a Symbol
	return parseAtom(head), nil
}

// parseNumber tries to parse a number from a string.
func parseNumber(s string) (number, error) {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return newNumber(0), fmt.Errorf("could not convert '%s' into a number", s)
	}
	return newNumber(val), nil
}

// parseSymbol parses a string into a Symbol.
func parseSymbol(s string) symbol {
	return newSymbol(s)
}

// parseAtom first tries to parse a token as a number, and if it fails, parses
// it as a Symbol.
func parseAtom(s string) Expression {
	maybeNum, err := parseNumber(s)
	if err != nil {
		return parseSymbol(s)
	}
	return maybeNum
}

// parseList parses a list of tokens into a List.
func parseList(tokens []string) list {
	l := newList()
	for i := 0; i < len(tokens); {
		switch t := tokens[i]; t {
		case ")":
			return l
		case "(":
			sublist := parseList(tokens[i+1:])
			l = l.push(sublist)
			i += sublist.length() + 2
		default:
			l = l.push(parseAtom(t))
			i++
		}
	}
	return l
}
