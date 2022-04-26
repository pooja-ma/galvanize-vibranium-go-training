package main

import "fmt"

func main() {
	// arrays
	arr := []string{"apple", "orange", "grapes"}
	arr2 := [5]int{4, 5, 6, 7, 8}
	fmt.Println(arr)
	fmt.Println(arr[1])
	fmt.Println(arr2)

	//simple for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//for range

	// k and l stores the value of arr
	// k store index number of individual string and
	// l store individual string of the given array
	for k, l := range arr {
		fmt.Println(k, l)
	}
}
