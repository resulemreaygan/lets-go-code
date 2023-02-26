/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func addNew(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	//Packages
	fmt.Println("My favorite number is", rand.Intn(10))

	//Imports
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	//Exported names
	// if it doesn't start with a capital letter, so they are not exported
	//fmt.Println(math.pi)

	fmt.Println(math.Pi)

	//Functions
	fmt.Println(add(42, 13))

	//Functions continued
	fmt.Println(addNew(42, 13))

	//Multiple results
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
