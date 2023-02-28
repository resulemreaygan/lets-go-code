/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	w := Vertex{1, 2}
	p := &w
	p.X = 1e9 // (*p).X not necessary
	fmt.Println(w)
}
