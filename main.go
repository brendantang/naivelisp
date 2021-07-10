package main

import (
	"bufio"
	"fmt"
	"github.com/brendantang/naivelisp/lisp"
	"log"
	"os"
)

func main() {
	log.Fatal(repl("naive-lisp:  "))
}

// Launch an interactive lisp prompt.
func repl(prompt string) error {
	env := lisp.StdEnv()
	for true {
		input, err := getInput()
		if err != nil {
			return err
		}
		exp, err := lisp.Parse(input)
		if err != nil {
			return err
		}
		val, newEnv, err := lisp.Eval(exp, env)
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
