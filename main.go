package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant work wih 42")
	}
	return arg + 3, nil
}

func main() {
	fmt.Println("Hwllow world ")
	//ERROR HANDLING IN GO

}
