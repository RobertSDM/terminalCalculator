package expression

import (
	"fmt"
	"regexp"
	"strconv"

	op "clc/cmd/operation"
	"clc/cmd/stack"
)

func isNumber(str string) (float64, bool) {
	value, err := strconv.ParseFloat(str, 64)
	return value, err == nil
}

var precedence = map[string]int{
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
	"(": 4,
	")": 4,
}

var alowedSymbolsRegex = `[0-9]+(?:\.[0-9]+)?|[-+/*().]`

func CleanExpression(expression string) []string {
	re := regexp.MustCompile(alowedSymbolsRegex)
	matches := re.FindAllString(expression, -1)

	if matches == nil {
		matches = make([]string, 0)
	}

	return matches
}

// Evaluates a Reverse Polish Notation
func SolveRPN(rpn []string) float64 {
	stk := stack.CreateStack[float64]()

	for _, exp := range rpn {
		if value, b := isNumber(exp); b {
			stk.Add(value)
		} else if stk.Len >= 2 {
			sv := stk.Pop()
			fv := stk.Pop()

			switch exp {
			case "+":
				stk.Add(op.Add(fv, sv))
			case "-":
				stk.Add(op.Subtract(fv, sv))
			case "*":
				stk.Add(op.Multiply(fv, sv))
			case "/":
				stk.Add(op.Divide(fv, sv))
			}
		}
	}

	return stk.Top()
}

func InfixToRPN(expression []string) []string {
	res := make([]string, 0)
	ops := stack.CreateStack[string]()

	var (
		addNegativeSignal bool
		addedNum          bool
	)
	for _, e := range expression {
		if _, b := isNumber(e); b {
			if addNegativeSignal {
				res = append(res, fmt.Sprintf("-%s", e))
			} else {
				res = append(res, e)
			}
			addedNum = true
		} else if e == "(" {
			ops.Add(e)
		} else if e == ")" {
			for ops.HasLen() && ops.Top() != "(" {
				res = append(res, ops.Pop())
			}
			ops.Pop()
		} else if e == "-" && (len(res) == 0 || !addedNum) {
			addNegativeSignal = true
		} else {
			for ops.HasLen() && precedence[ops.Top()] >= precedence[e] && ops.Top() != "(" {
				res = append(res, ops.Pop())
			}
			addedNum = false
			addNegativeSignal = false

			ops.Add(e)
		}
	}

	for ops.Len > 0 {
		res = append(res, ops.Pop())
	}

	return res
}

// Verify if all parentheses in the expression have a match
func MatchParentheses(expression []string) bool {
	stk := stack.CreateStack[string]()

	for _, e := range expression {
		if e == "(" {
			stk.Add(e)
		} else if e == ")" {
			if stk.Len == 0 {
				return false
			}

			stk.Pop()
		}
	}

	return stk.Len == 0
}

func ValidateExpression(expression []string) bool {
	var count int

	for _, e := range expression {
		if _, b := isNumber(e); b || e == "(" || e == ")" {
			count = 0
		} else {
			count += 1

			if count >= 2 {
				return false
			}
		}
	}

	return true
}
