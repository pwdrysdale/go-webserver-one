package main

import "fmt"

func aye() {
	var a int = 1
	var b int = a

	fmt.Println(a, b) // 1 1

	a = 2
	fmt.Println(a, b) // 2 1

	var c *int = &a
	fmt.Println(a, b, c) // 2 1 0xc0000aa008

	// so the value of c is the address of a in memory
	// note fmt.Println(&a) would also print 0xc0000aa008
	// as the & operator gets the address of a var	iable
	// so the * is pointer

	fmt.Println(*c) // 2

	a = 3
	fmt.Println(a, *c) // 3 3

	*c = 4
	fmt.Println(a, *c) // 4 4
}

// note that pointer arithmetic is not a thing (but is in some other languages)

func pointersInStructs() {
	type myStruct struct {
		a int
	}

	var b *myStruct = &myStruct{a: 1}
	fmt.Println(b) // &{1}
}

// slices and maps are also pointers, so be careful

func main() {
	aye()
	pointersInStructs()
}
