/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum // send sum to c
}

func main() {
	//Channels

	// ch <- v    // Send v to channel ch.
	// v := <-ch  // Receive from ch, and
	// assign value to v.

	// The data flows in the direction of the arrow.

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	// Like maps and slices, channels must be created before use

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
