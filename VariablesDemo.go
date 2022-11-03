package main

import (
	"fmt"
	"reflect"
)

func main() {

	var myInteger int
	myInteger = 1

	var myFloat float64
	myFloat = 3.1415

	var myBool bool
	myBool = true

	myShortIntVar := 1
	myShortFloatVar := 3.14

	fmt.Println(myInteger)
	fmt.Println(reflect.TypeOf(myInteger))

	fmt.Println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(myBool)
	fmt.Println(reflect.TypeOf(myBool))

	fmt.Println(myShortIntVar)
	fmt.Println(reflect.TypeOf(myShortIntVar))

	fmt.Println(myShortFloatVar)
	fmt.Println(reflect.TypeOf(myShortFloatVar))

	//Type conversion.
	var length float64 = 1.2
	var width int = 2
	fmt.Println("Area is: ", length*float64(width))

}
