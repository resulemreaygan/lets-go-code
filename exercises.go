/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"strconv"
)

func plus(a int, b int) int {
	return a + b
}

func plusplus(a int, b int, c int) int {
	return a + b + c
}

type Person struct {
	name      string
	age       int
	isMarried bool
}

type numbers struct {
	number1 float64
	number2 float64
}

func (n numbers) multiplication() float64 {
	return n.number1 * n.number2
}

type calc interface {
	multiplication() float64
}

func doCalc(c calc) {
	fmt.Println(c.multiplication())
}

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

	y := 10.7
	fmt.Println(3 + int(y))

	var tempString = "1"

	number, _ := strconv.Atoi(tempString)
	fmt.Println(4 + number)

	var abc int = 10

	word := strconv.Itoa(abc)
	fmt.Println(word)

	myArray := [...]string{"A", "B", "C", "D"}
	fmt.Println(myArray)

	for index, element := range myArray {
		fmt.Println(index, " ", element)
	}

	for _, element := range myArray {
		fmt.Println(element)
	}

	for index, _ := range myArray {
		fmt.Println(index)
	}

	for index := range myArray {
		fmt.Println(index)
	}

	fmt.Println(plus(3, 5))

	res := plusplus(5, 5, 5)
	fmt.Println(res)

	var p1 Person
	p1.name = "Resul Emre"
	p1.age = 26
	p1.isMarried = false

	fmt.Println(p1.name, " ", p1.age, " ", p1.isMarried)

	nbm := numbers{5, 5}

	doCalc(nbm)

	cityMap := map[int]string{
		58: "Sivas",
		34: "Istanbul",
		16: "Bursa",
		11: "Bilecik",
		06: "Ankara",
	}

	fmt.Println(cityMap)
}
