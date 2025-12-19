package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func basics() {
	//types of variable in golang
	var a = "variables"
	fmt.Println(a)

	var string = "string"
	fmt.Println(string)

	//variable declared wihtout a corresponding is declared a 0 valued variable

	// :- it is shorthand notation  to declare variables
	name := "kashish"

	age := 21

	fmt.Println(name, age)

	//constant in golang
	const n = 900000000000

	const d = 3e20 / n

	fmt.Println(math.Sin(n))
	fmt.Println(d)

}

func Forloop() {
	i := 1

	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	for i := range 3 {
		fmt.Println(i)
	}
}

func main() {
	basics()
	Forloop()
}
