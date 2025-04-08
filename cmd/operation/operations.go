package operation

func InvertSine(number int) int {
	return number * -1
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
