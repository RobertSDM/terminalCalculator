package main

import (
	"fmt"
	"os"

	exp "clc/cmd/expression"
)

func main() {
	if len(os.Args) < 2 {
		panic("please provide an expression")
	}

	rawExpression := os.Args[1]

	expression := exp.CleanExpression(rawExpression)

	if len(expression) == 0 {
		panic("no valid expression was found")
	}

	isValid := exp.ValidateExpression(expression)

	if !isValid {
		panic("invalid left-associative expression")
	}

	if !exp.MatchParentheses(expression) {
		panic("there are parentheses with no close matchs. Fix it")
	}

	rpn := exp.InfixToRPN(expression)
	fmt.Println(exp.SolveRPN(rpn))
}
