package expression

import (
	"strconv"
	"strings"

	op "clc/cmd/operation"
)

func IsNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

var precedence = map[rune]int{
	'*': 2,
	'/': 2,
	'+': 1,
	'-': 1,
	'(': 4,
	')': 4,
}

func SolveRPN(rpn []string) int {
	stack := []int{}

	for _, exp := range rpn {
		if IsNumber(exp) {
			re, _ := strconv.Atoi(exp)

			stack = append(stack, re)
		} else if len(stack) >= 2 {
			sv := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fv := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch exp {
			case "+":
				stack = append(stack, op.Add(fv, sv))
			case "-":
				stack = append(stack, op.Subtract(fv, sv))
			case "*":
				stack = append(stack, op.Multiply(fv, sv))
			case "/":
				stack = append(stack, op.Divide(fv, sv))
			}
		}
	}

	return stack[0]
}

// convert infix expression to RPN (Reverse Polish Notation)
func InfixToRPN(expression string) []string {
	res := make([]string, 0)
	ops := make([]rune, 0)

	var (
		value rune
		tmp   strings.Builder
	)
	for _, e := range expression {
		if !IsNumber(string(e)) && tmp.Len() > 0 {
			res = append(res, tmp.String())
			tmp.Reset()
		}

		if IsNumber(string(e)) {
			tmp.WriteRune(e)
		} else if e == '(' {
			ops = append(ops, e)
		} else if e == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				value, ops = ops[len(ops)-1], ops[:len(ops)-1]

				res = append(res, string(value))
			}
			ops = ops[:len(ops)-1]
		} else {
			for len(ops) > 0 && precedence[ops[len(ops)-1]] >= precedence[e] && ops[len(ops)-1] != '(' {
				value, ops = ops[len(ops)-1], ops[:len(ops)-1]
				res = append(res, string(value))
			}
			ops = append(ops, e)
		}
	}

	if tmp.Len() > 0 {
		res = append(res, tmp.String())
	}

	for len(ops) > 0 {
		value, ops = ops[len(ops)-1], ops[:len(ops)-1]
		res = append(res, string(value))
	}

	return res
}

func MatchParentheses(expression string) bool {
	stack := []rune{}

	for _, e := range expression {
		if e == '(' {
			stack = append(stack, e)
		} else if e == ')' {
			if len(stack) == 0 {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
