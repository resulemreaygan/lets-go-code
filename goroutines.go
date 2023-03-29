/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")

	//Goroutines run in the same address space, so access to shared memory must be synchronized.
	//The sync package provides useful primitives,
	//although you won't need them much in Go as there are other primitives.
}
