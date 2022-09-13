package main

import (
	"fmt"
)

func basicLoop() {
	// can use i++ or i += 1
	for i := 0; i < 10; i = i + 3 {
		fmt.Println(i)
	}
}

func twoIteratorLoop() {
	// you can use multiple iterators
	// note that the iterator is a statement, so
	// here you cannot use i++ or i += 1 and same with j
	for i, j := 0, 0; i < 10; i, j = i+3, j+3 {
		fmt.Println(i, j)
	}
}

func predeclareLoop() {
	// you can also predeclare the iterator
	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 4
	}
}

func usingBreak() {
	var i = 10
	for {
		fmt.Println(i)
		i = i - 3

		if i < 0 {
			break
		}
	}
}

// prints the odd values
func continueStatements() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func nestedLoop() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, j)
		}
	}
}

// using labels to break out of nested loops
func labels() {
	// you can use labels to break out of nested loops
	// note that the label is a statement, so
	// here you cannot use i++ or i += 1
outerLoop:
	for i := 0; i < 10; i = i + 3 {
		for j := 0; j < 10; j = j + 3 {
			fmt.Println(i, j)
			if i == 6 {
				break outerLoop
			}
		}
	}
}

func rangeLoop() {
	// you can also use range to iterate over arrays and slices
	// note that the range loop is a statement, so
	// here you cannot use i++ or i += 1

	// basically a setup for your js array methods
	for i, v := range []string{"one", "two", "three", "four", "five"} {
		fmt.Println(i, v)
	}
}

func main() {
	basicLoop()
	twoIteratorLoop()
	predeclareLoop()
	usingBreak()
	continueStatements()
	nestedLoop()
	labels()
	rangeLoop()
}
