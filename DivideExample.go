package main

import (
	"fmt"
	"log"
)

func main() {

	quotient, err := divide(5.6, 2.0)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Divisor : %.2f", quotient)
	}

}

func divide(dividend float64, divisor float64) (float64, error) {

	if divisor == 0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	} else {
		return dividend / divisor, nil
	}

}
