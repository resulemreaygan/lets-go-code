/*
Author: Resul Emre AYGAN
*/
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Number of Iteration %v, value = %v\n", i, z)
	}
	return z
}

func SqrtNew(x float64) float64 {
	z := 1.0
	var k float64

	for {
		z, k = z-(z*z-x)/(2*z), z

		fmt.Printf("z value = %v, k value = %v\n", z, k)

		if math.Abs(k-z) < 1e-6 {
			break
		}
	}
	return z
}

func main() {
	// Newton's method
	fmt.Println(Sqrt(2))

	guess := SqrtNew(2)
	expected := math.Sqrt(2)
	tolerance := math.Abs(guess - expected)

	fmt.Printf("Guess: %v, Expected: %v, Tolerance: %v", guess, expected, tolerance)
}
