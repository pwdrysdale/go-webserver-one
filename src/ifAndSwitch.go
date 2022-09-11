package main

import (
	"fmt"
	"math"
)

// if statements
// basic if statement
func ifStatement() {
	if true {
		fmt.Println("true")
	}
}

// if else if statement
func ifElseIfStatement(num int) {
	if num == 1 {
		fmt.Println("one")
	} else if num == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("not one or two")
	}
}

var statePopulations = map[string]int{
	"California": 39250017,
	"Texas":      27862596,
	"Florida":    20612439,
	"New York":   19745289,
}

func ifFlorida() {
	// one thing that is handy to know is that you can do some initializer work in the if statement
	// to do this, we do the work, then the iffy part is after the semicolon
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println("Florida baby! ", pop)
	}
}

// takes and input and returns true if it is truthy
func isTruthy(thing interface{}) bool {
	if thing != nil {
		return true
	}
	return false
}

// logicals include &&, ||, and ! (much like javascript)
// if you do use the or operator it will exit early (known as short circuiting?)
// and similar for the and

// floating point numbers can be a pain, its common practice to set some threshold
func checkEquality1(num1 float64) {
	if num1 == math.Pow(math.Sqrt(num1), 2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

func checkEquality2(num1 float64) {
	//  this is a common way to check for equality
	if math.Abs(num1-math.Pow(math.Sqrt(num1), 2)) < 0.0001 {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

// switch statements
// switch statements are like if statements, but they are more concise
// they can also be used to compare types
func switchStatement(num int) {
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}

// you can also do multiple cases
func switchMultipleCases(num int) {
	switch num {
	case 1, 2:
		fmt.Println("one or two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("not one or two")
	}
}

// you can also do ranges
// note that the logic here does not depend on one variable
func switchRange(num int) {
	switch {
	case num < 10:
		fmt.Println("less than 10")
		fallthrough // this will make it go to the next case as well and run it regardless of
		// whether it passes or not (so num = 8 would print both less than 10 and less than 20)
	case num < 20:
		fmt.Println("less than 20")
		break
		fmt.Println("this will never print")
	default:
		fmt.Println("not less than 10 or 20")
	}
}

// you can also do type switches
// this is a common way to check the type of an interface
// you can also use the type switch to get the type of the interface
func switchType(thing interface{}) {
	switch thing.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	ifFlorida()

	fmt.Println(isTruthy(1))
	fmt.Println(isTruthy(0))

	// these should both be equal
	checkEquality1(1.0)  // equal
	checkEquality1(1.01) // not equal

	checkEquality2(1.0)  // equal
	checkEquality2(1.01) // equal, as we built in some tolerance
}
