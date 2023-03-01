/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

func main() {
	//Arrays
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//Slices
	var s []int = primes[1:4]
	fmt.Println(s)
	// An array has a fixed size.
	// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

	//Slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	aNew := names[0:2]
	bNew := names[1:3]
	fmt.Println(aNew, bNew)

	bNew[0] = "XXX"
	fmt.Println(aNew, bNew)
	fmt.Println(names)
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.

	//Slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	sNew := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(sNew)

	//Slice defaults
	k := []int{2, 3, 5, 7, 11, 13}

	k = k[1:4]
	fmt.Println(k)

	k = k[:2]
	fmt.Println(k)

	k = k[1:]
	fmt.Println(k)

	//Slice length and capacity

	l := []int{2, 3, 5, 7, 11, 13}
	printSlice(l)

	// Slice the slice to give it zero length.
	l = l[:0]
	printSlice(l)

	// Extend its length.
	l = l[:4]
	printSlice(l)

	// Drop its first two values.
	l = l[2:]
	printSlice(l)

	//Nil slices
	var h []int
	fmt.Println(h, len(h), cap(h))

	if h == nil {
		fmt.Println("nil!")
	}

	//Creating a slice with make
	m := make([]int, 5)
	printSliceNew("m", m)

	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	// To specify a capacity, pass a third argument to make
	n := make([]int, 0, 5)
	printSliceNew("n", n)

	g := n[:2]
	printSliceNew("g", g)

	j := g[2:5]
	printSliceNew("j", j)
}

func printSlice(l []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(l), cap(l), l)
}

func printSliceNew(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(s), cap(x), x)
}
