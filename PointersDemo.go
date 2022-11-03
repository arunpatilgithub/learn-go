package main

import "fmt"

func main() {

	passByValueDemo()
	passByReferenceDemo()
	pointersExercise()
}

func passByValueDemo() {

	amount := 6
	doubleAmount(amount)
	//Here amount will NOT get doubled because go is pass my value by-default and not reference
	fmt.Println("Amount: ", amount)
}

func doubleAmount(num int) {
	num *= 2
}

func passByReferenceDemo() {
	amount := 6

	doubleAmountWithPointers(&amount)
	//This time, amount should get doubled
	fmt.Println("Amount: ", amount)

}

func doubleAmountWithPointers(amountVal *int) {
	*amountVal *= 2
}

/*


// negate takes a boolean value and returns its
// opposite. E.g.: negate(false) returns true.
// But we WANT this function to accept a POINTER
// to a boolean value, and update the value at
// the pointer to its opposite. Once this change
// is made, the function doesn't need to return
// anything.


func negate(myBoolean bool) bool {
	return !myBoolean
}

func pointersExercise() {
	truth := true
	// Change this to pass a pointer.
	negate(truth)
	// Prints "true", but we want "false".
	fmt.Println(truth)
	lies := false
	// Change this to pass a pointer.
	negate(lies)
	// Prints "false", but we want "true".
	fmt.Println(lies)
}


*/
func pointersExercise() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
