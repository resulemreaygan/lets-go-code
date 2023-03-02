/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	firstNumber, secondNumber := 0, 1

	return func() int {
		tempNumber := firstNumber
		firstNumber, secondNumber = secondNumber, tempNumber+secondNumber
		return tempNumber
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
