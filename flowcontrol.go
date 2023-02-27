/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

	//Switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows..

		fmt.Printf("%s.\n", os)
	}

	//Switch evaluation order
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In tow days.")
	default:
		fmt.Println("Too far away.")
	}

	//Switch with no condition
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	//Defer
	defer fmt.Println("world")

	fmt.Println("Hello")

	//Stacking defers
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		//resource leak in loop
	}

	fmt.Println("done")
}
