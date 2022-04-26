package main

import (
	"fmt"
	"reflect"
)

var booltype bool

var i, j int = 1, 2 //variables initialized.

func main() {
	var i int
	var str string

	//default - zerovalue
	//0 for numeric types,
	//false for the boolean type, and
	//"" (the empty string) for strings.
	fmt.Println(i, booltype)
	fmt.Println(str)

	//short variable declarations - if a variable is initiliazed , type can be omiited
	k := 3
	var jk int = 3
	// This should be inside a function
	fmt.Println(reflect.TypeOf(k))
	fmt.Println(reflect.DeepEqual(jk, k))
}

//int  int8  int16  int32  int64 uint uint8 uint16 uint32 uint64 uintptr
//byte - alias for uint8

//rune - alias for int32  represents a Unicode code point

//float32 float64

//complex64 complex128
