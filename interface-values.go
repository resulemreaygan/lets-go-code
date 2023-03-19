/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"math"
)

type I2 interface {
	M2()
}

type T2 struct {
	S string
}

func (t *T2) M2() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M2() {
	fmt.Println(f)
}

func main() {
	var i I2

	i = &T2{"Hello"}
	describe(i)
	i.M2()

	i = F(math.Pi)
	describe(i)
	i.M2()

	var j I2

	var t *T2
	j = t
	describe(j)
	j.M2()

	j = &T2{"hello"}
	describe(j)

	j.M2()
}

func describe(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}
