/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

func (v *Vertex3) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex3) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	//Methods and pointer indirection
	v := Vertex3{3, 4}
	v.Scale2(2)
	ScaleFunc(&v, 10)

	p := &Vertex3{4, 3}
	p.Scale2(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	//Methods and pointer indirection (2)
	q := Vertex3{3, 4}
	fmt.Println(q.Abs2())
	fmt.Println(AbsFunc(q))

	w := &Vertex3{4, 3}
	fmt.Println(w.Abs2())
	fmt.Println(AbsFunc(*w))
}
