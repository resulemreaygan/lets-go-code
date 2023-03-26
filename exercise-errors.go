/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

//func Sqrt(x float64) (float64, error) {
//	return 0, nil
//}

func Sqrt2(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	startVal := 1.0
	var tempVal float64

	for {
		startVal, tempVal = startVal-(startVal*startVal-x)/(2*startVal), startVal

		if math.Abs(tempVal-startVal) < 1e-6 {
			return startVal, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}
