// https://www.youtube.com/watch?v=YS4e4q9oBaU
package main

import (
	"fmt"
	"strconv"
)

// common to declare many variables at the module level
var foo string = "foo"
var bar string = "bar"

// you can short hand this a bit too

var (
	baz string = "baz"
	qux string = "qux"
)

// scoping
var (
	number int = 3
)

// variables with uppercase are available outside of the function
var Uppercase string = "uppercase"


// function that takes an argumet that is a string or a number
func printArg(arg interface{}) {
	fmt.Println(arg)
}

func print(number interface{} ) {
		fmt.Println(number)
	}

func longPrint(inp int) {
	// pull off the values and the type then set them
	fmt.Printf("%v, %T \n", inp , inp)
}

func main() {
	// shadowing example
	print("Shadowing")
	print(number)
	// in shadowing you can redeclare a variable, and the innermost one will be used


	// variables can be declared in three ways:
	// 1. declare the name and the type, then set a value
	var number int 
	number = 10

	print("Number: ")
	print(number)

	// 2. declare the name and the type, then set a value
	var anotherNumber int = 20
	print("Another Number: ")
	print(anotherNumber)
	
// 3. allow type inference
// note that this is not allowed outisde of a function
 	yetAnother := 4

	print("Yet Another: ")
	print(yetAnother)

	longPrint(yetAnother)
	print(qux)

	// convert types 
	const intey int = 10
	const floaty float64 = float64(intey)
	print(floaty)
	
	// to convert strings you need to use the strconv package	
	const intey2 int = 10
	// integer to a (string)
	var stringy string = strconv.Itoa(intey2)
	printArg(	stringy)

	const someNumber string = "10"
	// as a string
	stringy2, err := strconv.Atoi(someNumber)
	printArg(	stringy2 )
	printArg(	err )

}