package main

import (
	"errors"
	"fmt"
	"math"
)

// Go programs express error state with 'error' values.
// The error type is a built-in interface similar to fmt.Stringer:

type error interface {
	Error() string
}

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}
func main() {
	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = Sqrt(9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
