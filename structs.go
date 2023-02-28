/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	d  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	w := Vertex{1, 2}
	p := &w
	p.X = 1e9 // (*p).X not necessary
	fmt.Println(w)

	fmt.Println(v1, d, v2, v3)
}
