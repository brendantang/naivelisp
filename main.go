package main

import (
	"bufio"
	"fmt"
	"github.com/brendantang/naivelisp/expression"
	"github.com/brendantang/naivelisp/parser"
	"log"
	"os"
)

func main() {
	log.Fatal(repl("naive-lisp:  "))
}

// Launch an interactive lisp prompt.
func repl(prompt string) error {
	var env expression.Environment
	for true {
		input, err := getInput()
		if err != nil {
			return err
		}
		expressions, err := parser.Parse(input)
		if err != nil {
			return err
		}
		val, newEnv, err := expression.Eval(expressions, env)
		if err != nil {
			return err
		}
		if val != nil {
			fmt.Println(val.String())
		}
		env = newEnv
	}
	return nil
}

// Read a line from stdin.
func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}
