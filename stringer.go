/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

type Person2 struct {
	Name string
	Age  int
}

func (p Person2) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person2{"Resul Emre AYGAN", 26}
	b := Person2{"R. Emre AYGAN", 9001}

	fmt.Println(a, b)
}
