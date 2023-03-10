/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"strings"
)

var powArray = []int{1, 2, 4, 8, 16, 32, 64, 128}

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

	//Slices of slices

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//Appending to a slice
	var f []int
	printSlice(f)

	// append works on nil slices.
	f = append(f, 0)
	printSlice(f)

	// The slice grows as needed.
	f = append(f, 1)
	printSlice(f)

	// We can add more than one element at a time.
	f = append(f, 2, 3, 4)
	printSlice(f)

	// If the backing array of f is too small to fit all the given values a bigger array will be allocated.
	// The returned slice will point to the newly allocated array.

	//Range
	for i, v := range powArray {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//Range continued
	powNew2 := make([]int, 10)

	// If you only want the index, you can omit the second variable.
	for i := range powNew2 {
		powNew2[i] = 1 << uint(i) // == 2**i
	}

	// You can skip the index or value by assigning to _.
	for _, value := range powNew2 {
		fmt.Printf("%d\n", value)
	}
}

func printSlice(l []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(l), cap(l), l)
}

func printSliceNew(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(s), cap(x), x)
}
