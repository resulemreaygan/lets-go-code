/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
	//return v
}

func powNew(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	//For
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//For continued
	sumNew := 1

	for sumNew < 1000 {
		sumNew += sumNew
	}
	fmt.Println(sumNew)

	//For is Go's "while"
	sumNew2 := 1

	for sumNew2 < 1000 {
		sumNew2 += sumNew2
	}
	fmt.Println(sumNew2)

	//Forever
	// infinite loop
	//for {
	//}

	//If
	fmt.Println(sqrt(2), sqrt(-4))

	//If with a short statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//If and else
	fmt.Println(
		powNew(3, 2, 10),
		powNew(3, 3, 20),
	)
}
