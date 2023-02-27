/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"math"
	"math/cmplx"
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

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var iNew, jNew int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
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

	//Named return values
	//Naked return statements
	fmt.Println(split(17))

	//Variables
	var i int
	fmt.Println(i, c, python, java)

	//Variables with initializers
	var cNew, pythonNew, javaNew = true, false, "no!"
	fmt.Println(iNew, jNew, cNew, pythonNew, javaNew)

	//Short variable declarations
	//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	var iNew2, jNew2 int = 1, 2
	k := 3
	cNew2, pythonNew2, javaNew2 := true, false, "no!"

	fmt.Println(iNew2, jNew2, k, cNew2, pythonNew2, javaNew2)

	//Basic types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//Zero values
	var iNew3 int
	var f float64
	var bNew bool
	var s string
	fmt.Printf("%v %v %v %q\n", iNew3, f, bNew, s)

	//Type conversions
	var xNew, yNew int = 3, 4
	var fNew float64 = math.Sqrt(float64(xNew*xNew + yNew*yNew))
	var zNew uint = uint(fNew)
	fmt.Println(xNew, yNew, zNew)

	//Type inference
	v := 42 //change me!
	fmt.Printf("v is of type %T\n", v)

	//Constants
	//cannot be declared using the := syntax
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	//Numeric Constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big))
}
