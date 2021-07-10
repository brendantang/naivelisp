package lisp

import (
	"testing"
)

func TestEval(t *testing.T) {
	for _, c := range evalTestCases {
		var xs []Expression
		for i, line := range c.srcLines {
			exp, err := Parse(line)
			if err != nil {
				t.Fatalf("ERROR: %s\nInput line %d: %s\nGot error from Parse: %e", c.description, i, line, err)
			}
			xs = append(xs, exp)
		}
		var gotExp Expression
		env := StdEnv()
		for i, exp := range xs {
			newExp, newEnv, err := Eval(exp, env)
			if err != nil {
				t.Fatalf("ERROR: %s\nInput line %d: %s\nGot error from Parse: %e", c.description, i, exp.String(), err)
			}
			gotExp, env = newExp, newEnv
		}
		if c.wantExp != gotExp.String() {
			t.Fatalf(
				"FAIL: %s\nInput: %s\nExpected: %s\nGot: %s\n",
				c.description,
				c.srcLines,
				c.wantExp,
				gotExp.String(),
			)
		}
	}
}

var evalTestCases = []struct {
	description string
	srcLines    []string
	wantExp     string
}{
	{
		description: "number evaluates to itself",
		srcLines:    []string{"3"},
		wantExp:     "3",
	},
	{
		description: "primitive procedure call, addition",
		srcLines:    []string{"(+ 3 45)"},
		wantExp:     "48",
	},
	{
		description: "primitive procedure call, subtraction",
		srcLines:    []string{"(- 300 100)"},
		wantExp:     "200",
	},
	{
		description: "nested procedure call",
		srcLines:    []string{"(+ 10 (- 300 100))"},
		wantExp:     "210",
	},
	{
		description: "simple definition",
		srcLines:    []string{"(define x 333)", "x"},
		wantExp:     "333",
	},
	{
		description: "proc definition",
		srcLines:    []string{"(define (sq x) (* x x))", "(sq 5)"},
		wantExp:     "25",
	},
}
