package main

import (
	"fmt"
)

//  Booleans
var (
	boolz bool = true
	boolz2 bool = false
)

func flipBool(inp bool) bool{
	return !inp
}

//  Integers

// 8 bit
//  range from -128 to 127
var (
	int8z int8 = 127
	int8z2 int8 = -128
)

// 16 bit
//  range from -32768 to 32767
var (
	int16z int16 = 32767
	int16z2 int16 = -32768
)

// 32 bit
//  range from -2147483648 to 2147483647
var (
	int32z int32 = 2147483647
	int32z2 int32 = -2147483648
)

// 64 bit
//  range from -9223372036854775808 to 9223372036854775807
var (
	int64z int64 = 9223372036854775807
	int64z2 int64 = -9223372036854775808
)

//  Unsigned Integers
// Unsigned integers are positive numbers only

// 8 bit
//  range from 0 to 255
var (
	uint8z uint8 = 255
	uint8z2 uint8 = 0
)
//  etc...
// but there is no unsigned int64

//  Floats
//  Floats are numbers with a decimal point
var (
	float32z float32 = 3.14
	float32z2 float32 = 6.28e-10
	float32z3 float32 = 1.2E4
)

//  Complex Numbers
//  Complex numbers are numbers with a real and imaginary part
var (
	complex64z complex64 = 1 + 2i
	complex64z2 complex64 = 1.0 + 2.0i
	complex128z complex128 = 1 + 2i
)

//  can use complex(real, imaginary) to create a complex number
//  can use real(complex) to get the real part of a complex number
//  can use imag(complex) to get the imaginary part of a complex number

//  Strings
//  Strings are a sequence of bytes
// They are immutable, so you can't change them
var (
	stringz string = "This is a string"
)

//  Rune 
//  A rune is an alias for int32 and is used to represent a Unicode code point
var (
	runez rune = 'a' // declared with a single quote - runez has the value 97
)



func main(){
	fmt.Println("Boolz: ", boolz)
	fmt.Println("Boolz2: ", boolz2)
	fmt.Println("Flipped Boolz: ", flipBool(boolz))

	a := 10
	b := 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // 3, as when you divide two integers you get an integer
	fmt.Println(a % b)

	//  You probably want to do type conversion to keep the mathematics safe, as mixing types can cause errors
	var c int64 = 10
	var d int8 = 3

	// fmt.Println(c + d)   throws
	//  cannot use d (type int8) as type int64 in addition

	fmt.Println(c + int64(d)) // 13

	// Bitwise operators
	// I'll skip these here, but they are useful for low level programming

	// Float ops
	var e float32 = 10.2
	var f float32 = 3.7

	fmt.Println(e + f)
	fmt.Println(e - f)
	fmt.Println(e * f)
	fmt.Println(e / f)
	//  there is no modulus operator for floats

	//  byte slice operations
	someString  := "This is a string"
	asBytes := []byte(someString)
	fmt.Println(asBytes)

	// concatenation
	//  you can concatenate strings with the + operator
	longString := "This is a long string" + " and this is another string"
	fmt.Println(longString)
}

