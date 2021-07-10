package lisp

import (
	"bufio"
	"fmt"
	"os"
)

// Interactive launches an interactive lisp prompt.
func Interactive(prompt string) error {
	env := StdEnv()
	for true {
		input, err := getInput(prompt)
		if err != nil {
			return err
		}
		val, newEnv, err := interpretLine(input, env)
		if val != nil {
			fmt.Println("=> " + val.String())
		}
		env = newEnv
	}
	return nil
}

// read a line from stdin.
func getInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func interpretLine(line string, env Environment) (exp Expression, newEnv Environment, err error) {
	exp, err = Parse(line)
	if err != nil {
		return
	}
	exp, newEnv, err = Eval(exp, env)
	if err != nil {
		return
	}
	return
}
