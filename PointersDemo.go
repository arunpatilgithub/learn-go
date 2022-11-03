package main

import "fmt"

func negate(myBoolean *bool) {
	*myBoolean = !(*myBoolean)
}

func main() {
	truth := true
	// Change this to pass a pointer.
	negate(&truth)
	// Prints "true", but we want "false".
	fmt.Println(truth)
	lies := false
	// Change this to pass a pointer.
	negate(&lies)
	// Prints "false", but we want "true".
	fmt.Println(lies)
}
