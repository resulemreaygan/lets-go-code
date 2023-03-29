/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
)

func fibonacci2(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	//Range and Close

	// v, ok := <-ch
	c := make(chan int, 10)

	go fibonacci2(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}

	// Only the sender should close a channel, never the receiver.
	// Sending on a closed channel will cause a panic.

	// Channels aren't like files; you don't usually need to close them.
	// Closing is only necessary when the receiver must be told there are no more values coming,
	// such as to terminate a range loop.
}
