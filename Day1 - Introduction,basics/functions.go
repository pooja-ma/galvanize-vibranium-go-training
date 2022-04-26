package main

import "fmt"

func main() {
	fmt.Println(add1(10, 20))

	fmt.Println(add2(15, 30))
	arr := []string{"apple", "orange", "grapes"}
	arr2 := []string{"apple1", "orange1", "grapes1"}
	a, b := swap("hello", "world")
	a1, b1 := swap1(arr, arr2)
	fmt.Println(a, b)
	fmt.Println(a1, b1)
}

// A function can take zero or more arguments.
// notice that the type comes after the variable name.

func add1(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.
func add2(x, y int) int {
	return x + y
}

// A function can return any number of results.
// The swap function returns two strings.
func swap(x, y string) (string, string) {
	return y, x
}

func swap1(x, y []string) ([]string, []string) {
	return y, x
}
