/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

func main() {
	// Arrays
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	var s []int = primes[1:4]
	fmt.Println(s)
	// An array has a fixed size.
	// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
}
