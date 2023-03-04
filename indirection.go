/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

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

func main() {
	//Methods and pointer indirection
	v := Vertex3{3, 4}
	v.Scale2(2)
	ScaleFunc(&v, 10)

	p := &Vertex3{4, 3}
	p.Scale2(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
