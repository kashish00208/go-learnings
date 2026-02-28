package main

import (
	"fmt"
)


func printFor() {
	fmt.Println("hello world")
	fmt.Println("go" + "lang")
	fmt.Print("1+1 = ", 1+1)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(constantVar)
	k := 19
	for j := 1; j <= 10; j++ {
		fmt.Println(k, " * ", j, " ", k*j)
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := range 100 {
		fmt.Println(i)
	}
}

func main() {
	var a [5]int
	a[3] = 12

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	fmt.Println(len((b)))

	arr := [...]int(100,3:300,23)

	fmt.Println(a)
}
