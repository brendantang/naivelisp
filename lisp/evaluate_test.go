package lisp

import (
	"testing"
)

func TestEval(t *testing.T) {
	for _, c := range evalTestCases {
		gotExp, gotEnv, err := Eval(c.haveExp, c.haveEnv)
		if err != nil {
			t.Fatalf("Got error from Eval: %e", err)
		}
		if !(c.wantExp.Equal(gotExp) && c.wantEnv.Equal(gotEnv)) {
			t.Fatalf(
				"FAIL: %s\nInput: %v\n%v\n\nExpected: %v\n%v\nGot: %v\n%v\n",
				c.description,
				c.haveExp,
				c.haveEnv,
				c.wantExp,
				c.wantEnv,
				gotExp,
				gotEnv,
			)
		}
	}
}

var evalTestCases = []struct {
	description string
	haveExp     Expression
	haveEnv     Environment
	wantExp     Expression
	wantEnv     Environment
}{
	{
		description: "number",
		haveExp:     NewNumber(3),
		haveEnv:     StdEnv(),
		wantExp:     NewNumber(3),
		wantEnv:     StdEnv(),
	},
}
