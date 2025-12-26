package main

import (
	"fmt"
)

func main() {
	fmt.Println("ARRAY IMPLENTATION")

	var a [5]int

	a[2] = 1

	fmt.Println(a[4])

	for i := 0; i < len(a); i++ {
		a[i] = i
		//fmt.Println("a[", i, "]", "is", i)
	}

	nums := [5]int{1, 9, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		//println(nums[i])
	}

	b := [...]int{1, 2, 3, 4, 4, 5, 5, 5}

	for i := 0; i < len(b); i++ {
		//fmt.Println(b[i])
	}

	//2D Array implementation
	var arr [2][2]int

	//var array[2][2]int{{1,2},{2,3},{3,4}}

	for i := range 2 {
		for j := range 2 {
			arr[i][j] = i + j
		}
	}

	for i := range 2 {
		for j := range 2 {
			fmt.Println(i, j)
		}
	}

}
