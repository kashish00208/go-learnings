package main

import (
	"fmt"
	"time"
)

// Constant declaration
const s string = "Constant"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Constant value:", s)

	// VARIABLES & BASIC OPERATIONS
	var name = "Kashish"
	var x, y int = 1, 2
	var isAwesome = true
	var uninitialized int

	fmt.Println("\n--- Variables ---")
	fmt.Println("Name:", name)
	fmt.Println("x:", x, "y:", y)
	fmt.Println("isAwesome:", isAwesome)
	fmt.Println("Uninitialized (default zero value):", uninitialized)

	// Short variable declaration
	person := "Kashish"
	age := 21
	suffix := "Ji"

	fmt.Println(person, age, suffix)

	// CONSTANTS & MATH EXAMPLE
	const n = 500
	const d = 3e20 / n
	fmt.Println("Value of d (3e20 / n):", d)
	fmt.Println("Converted to int64:", int64(d))

	// LOOPS

	// 1. While-style loop
	i := 1
	for i <= 3 {
		fmt.Println("While-style loop:", i)
		i++
	}

	// 2. Classic for loop
	for j := 0; j < 3; j++ {
		fmt.Println("Classic for loop:", j)
	}

	// 3. Range over integer (Go 1.22+ feature)
	for k := range 3 {
		fmt.Println("Range loop (0..2):", k)
	}

	// 4. Infinite loop (with break)
	for {
		fmt.Println("Infinite loop once")
		break
	}

	// 5. Loop with continue (prints odd numbers only)
	for n := range 10 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Odd number:", n)
	}

	// TIME EXAMPLE
	fmt.Println("Current time:", time.Now())

	fmt.Println("\nProgram executed successfully ✅")

	//SWITCH STATEMENTS

	switch i {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a weekend")

	default:
		fmt.Println("Yes you have to work")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Workout first and the work")
	default:
		fmt.Println("Yes it is the time of work go work hard you smart bitch")
	}
	fmt.Println("BHAGWAN JI KISI KO BHEJ DO MUJE GURDWARE JANA HAI")

	var a [7]int
	fmt.Println(a[6])
}
