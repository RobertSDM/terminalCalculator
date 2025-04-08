package expression

import (
	"regexp"
	"strconv"

	op "clc/cmd/operation"
)

func IsNumber(str string) bool {
	return regexp.MustCompile(`\d`).MatchString(str)
}

var precedence = map[string]int{
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
	"(": 4,
	")": 4,
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
	res := []string{}
	ops := []string{}

	var value string
	tmp := ""
	for _, e := range expression {
		if IsNumber(string(e)) {
			tmp += string(e)
		} else if e == '(' {
			if tmp != "" {
				res = append(res, string(tmp))
				tmp = ""
			}

			ops = append(ops, string(e))
		} else if e == ')' {
			if tmp != "" {
				res = append(res, string(tmp))
				tmp = ""
			}

			for len(ops) > 0 && string(ops[len(ops)-1]) != "(" {
				value, ops = ops[len(ops)-1], ops[:len(ops)-1]

				res = append(res, value)
			}
			ops = ops[:len(ops)-1]
		} else {
			if tmp != "" {
				res = append(res, string(tmp))
				tmp = ""
			}

			for len(ops) > 0 && precedence[string(ops[len(ops)-1])] >= precedence[string(e)] && ops[len(ops)-1] != "(" {
				value, ops = ops[len(ops)-1], ops[:len(ops)-1]
				res = append(res, value)
			}
			ops = append(ops, string(e))
		}
	}
	if tmp != "" {
		res = append(res, tmp)
	}

	for len(ops) > 0 {
		value, ops = ops[len(ops)-1], ops[:len(ops)-1]
		res = append(res, value)
	}

	return res
}

func MatchParenteses(expression string) bool {
	stack := []rune{}

	for _, exp := range expression {
		if exp == '(' {
			stack = append(stack, exp)
		} else if len(stack) > 0 && exp == ')' {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
