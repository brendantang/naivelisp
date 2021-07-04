package main

import (
	"bufio"
	"fmt"
	"github.com/brendantang/naive-lisp/interpreter"
	"github.com/brendantang/naive-lisp/parser"
	"log"
	"os"
)

func main() {
	log.Fatal(Start("naive-lisp:  "))
}

// Start launches an interactive lisp prompt.
func Start(prompt string) error {
	for true {
		input, err := GetInput()
		if err != nil {
			return err
		}
		val, err := interpreter.Eval(parser.Parse(input))
		if err != nil {
			return err
		}
		if val != nil {
			fmt.Println(LispString(val))
		}

	}
	return nil
}

// GetInput reads a line from stdin.
func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

// LispString converts a Lisp value into a readable string.
func LispString(interpreter.Expression) string {
	return "Implement lisp string"
}
