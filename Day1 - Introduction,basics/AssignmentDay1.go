package main

import (
	"fmt"
	"errors"
)

//Function which returns quotient if the given numbers are valid else returns err
func div(divident float64, divisor float64) (float64, error){
	if(divident == 0 || divisor == 0){
		return 0, errors.New("Error: Divident or Divisor cannot be zero")
	}
	return divident/divisor, nil
}

//Function which returns only odd indexed items for the given array 
func printOddIndexItemsOfArray(array []string) []string{
	oddIndexArr := []string{}
    for i := 1; i < len(array); i += 2 {
		oddIndexArr = append(oddIndexArr, array[i])	
    }
	return oddIndexArr
}

//main function
func main(){
	quotient, err := div(6,3)
	fmt.Print("Division of 6 by 3---> ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Quotient: ",quotient)
	}
	fmt.Print("Division of 0 by 3---> ")
	quotient1, err1 := div(0,3)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Quotient: ",quotient1)
	}
	fmt.Println("-----------------------------")
	array := []string{"1","2","3","4","5"}
	fmt.Println("Array: ", array)
	oddindxArr := printOddIndexItemsOfArray(array)
	fmt.Println("Odd indexed Array: ",oddindxArr)
}