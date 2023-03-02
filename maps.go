/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

type VertexNew struct {
	Lat, Long float64
}

var m map[string]VertexNew

func main() {
	//Maps
	m = make(map[string]VertexNew)
	m["Bell Labs"] = VertexNew{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
}
