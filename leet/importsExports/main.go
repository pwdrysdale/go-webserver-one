package main

import (
	"fmt"
	"importsExports/subpackage"
)

func main() {
	fmt.Println("Hello, world!")   // prints Hello, world!
	fmt.Println("Hello,", Exports) // prints Hello, Peter
	fmt.Println("Hello,", fib(10)) // prints Hello, 55

	fmt.Println(subpackage.Sub)
	subpackage.SomeFunction()
	subpackage.Subpackage()
	fmt.Println(subpackage.Stringz)
}
