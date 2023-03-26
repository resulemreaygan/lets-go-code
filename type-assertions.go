/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

func main() {
	//Type assertions
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	//fmt.Println(f)

	var j interface{} = 56.2

	g, ok := j.(float64)
	fmt.Println(g, ok)

	h, ok := j.(float32)
	fmt.Println(h, ok)

	k, ok := j.(int32)
	fmt.Println(k, ok)

}
