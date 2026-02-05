package main

import "fmt"

type Operation struct {
	A      int
	B      int
	Result int
	Op     string
}

var history []Operation

func calculate(a, b int, op string) (int, error) {
	var res int

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		return 0, fmt.Errorf("Unknown Operator")
	}

	history = append(history, Operation{a, b, res, op})
	return res, nil
}
func showHistory() {
	fmt.Println("Operation History:")
	for i, op := range history {
		fmt.Printf("%d. %d %s %d = %d\n", i+1, op.A, op.Op, op.B, op.Result)
	}
}

func main() {
	calculate(10, 5, "+")
	calculate(20, 4, "-")
	calculate(3, 7, "*")
	showHistory()
}
