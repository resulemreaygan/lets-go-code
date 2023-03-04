/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsNew(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) AbsNew2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleNew(v *Vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	//Methods
	v := Vertex2{3, 4}
	fmt.Println(v.Abs())

	//Methods are functions
	fmt.Println(AbsNew(v))

	//Methods continued
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.AbsNew2())

	//Pointer receivers
	v.Scale(10)
	fmt.Println(v.Abs())

	//Pointers and functions
	vNew := Vertex2{3, 4}
	ScaleNew(&vNew, 10)
	fmt.Println(AbsNew(vNew))
}
