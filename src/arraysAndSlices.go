package main

import "fmt"

// Arrays
// Arrays are a fixed length list of elements of a single type

var arr [5]int // array of 5 integers
var arr2 = [...]int{1, 2, 3, 4, 5} // array of 5 integers
var arr3 = []string{"a", "b", "c"} // slice of strings (looks like an array though!)

// Slices
// Slices are a variable length list of elements of a single type
var slice = []int{1, 2, 3, 4, 5} // slice of 5 integers




func main() {
	fmt.Println(arr)
	fmt.Println(arr2)

	fmt.Println(len(arr3))

	arr4 := []int{1, 2, 3, 4, 5} // slice of integers
	arr4[0] = arr4[1]
	fmt.Println(arr4)

	slice = append(slice, 6)
	fmt.Println(slice) // works [1 2 3 4 5 6]

	// slices kind of use pointers by default
	var a = []int{1, 2, 3, 4, 5}
	var b = a
	b[0] = 100

	fmt.Println(a) // [100 2 3 4 5]
	fmt.Println(b) // [100 2 3 4 5]

	// to copy a slice you need to use the copy function
	var c = make([]int, 5)
	copy(c, a)
	c[0] = 200
	fmt.Println(a) // [100 2 3 4 5]
	fmt.Println(c) // [200 2 3 4 5]

	// you can also use the append function to copy a slice
	var d = append([]int{}, a...)
	d[0] = 300
	fmt.Println(a) // [100 2 3 4 5]
	fmt.Println(d) // [300 2 3 4 5]

	// you can create slices from slices using indices
	var e = a[1:3]
	fmt.Println(e) // [2 3]
	var f = a[1:]
	fmt.Println(f) // [2 3 4 5]
	var g = a[:3]
	fmt.Println(g) // [100 2 3]

}

//  The capacity is the length of the underlying array, and the length is the number of elements referred to by the slice