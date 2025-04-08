package main

import (
	"os"

	exp "clc/cmd/expression"
)

func main() {
	if len(os.Args) < 2 {
		panic("please provide an expression")
	}

	expression := os.Args[1]

	if !exp.MatchParenteses(expression) {
		panic("there are parenteses with no close matchs. Fix it")
	}

	rpn := exp.InfixToRPN(expression)
	println(exp.SolveRPN(rpn))
}
