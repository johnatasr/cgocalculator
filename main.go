package main

// #include "calculator.c"
import "C"

import (
	"fmt"
)

func addition(x, y float64) float64 {
	return float64(C.addition(C.double(x), C.double(y)))
}

func subtract(x, y float64) float64 {
	return float64(C.subtract(C.double(x), C.double(y)))
}

func multiply(x, y float64) float64 {
	return float64(C.multiply(C.double(x), C.double(y)))
}

func divide(x, y float64) float64 {
	return float64(C.divide(C.double(x), C.double(y)))
}

func main() {
	x := 10.0
	y := 5.0

	fmt.Printf("Addition: %.2f\n", addition(x, y))
	fmt.Printf("Subtraction: %.2f\n", subtract(x, y))
	fmt.Printf("Multiplication: %.2f\n", multiply(x, y))
	fmt.Printf("Division: %.2f\n", divide(x, y))
}
