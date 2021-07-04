package main

import (
	"fmt"
)

// Values are
type Value interface {
	toString() string
}

type Symbol string

func (s Symbol) toString() string {
	return fmt.Sprint(s)
}

type Number float64

func (n Number) toString() string {
	return fmt.Sprintf("%v", n)

}

type List []Value

func (l List) toString() string {
	return "a list"
}
