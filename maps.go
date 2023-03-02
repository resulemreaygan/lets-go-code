/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

type VertexNew struct {
	Lat, Long float64
}

var m map[string]VertexNew

var n = map[string]VertexNew{
	"Bell Labs": VertexNew{40.68433, -74.39967},
	"Google":    VertexNew{37.42202, -122.08408},
}

var k = map[string]VertexNew{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	//Maps
	m = make(map[string]VertexNew)
	m["Bell Labs"] = VertexNew{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])

	//Map literals
	fmt.Println(n)

	//Map literals continued
	fmt.Println(k)
}
