package main

import (
	"fmt"
	"math"
	"time"
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

func ifelseSwitch() {
	fmt.Println("Complete golang basics")

	// if else practice in golang

	if 4%2 == 0 {
		fmt.Println("It works")
	}

	name := "kashish"

	if name == "kashish" {
		fmt.Println("Yes")
	}

	var ans int = 1234 * 2

	if ans%2 == 0 {
		fmt.Println("yes is divisible by 2")
	}

	const answer int = 23

	if ans == 23 {
		fmt.Println(answer + 23)
	}

	//Switch cases in golang

	i := 3

	fmt.Println("i as", i)

	switch i {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend ")
	default:
		fmt.Println("It weekday you have to work")
	}

	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("Do not push code in prod its friday")
	}
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
