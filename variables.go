/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

func main() {
	var name string

	name = "Resul Emre"

	fmt.Println(name)

	var nameTemp = "Resul Emre"

	fmt.Println(nameTemp)

	numberTemp := 123
	fmt.Println(numberTemp)

	numberTemp2, booleanTemp, stringTemp := 123, true, "REA"

	fmt.Println(numberTemp2, booleanTemp, stringTemp)
}
