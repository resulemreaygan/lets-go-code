/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	//Default Selection

	// select {
	// case i := <-c:
	//    // use i
	// default:
	//    // receiving from c would block
	// }
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
