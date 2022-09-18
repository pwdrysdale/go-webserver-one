package main

import (
	"fmt"
	"strconv"
	"strings"
)

// declaring basic functions
func add(x int, y int) int {
	return x + y
}

// declaring functions with multiple return values
// here we take two arguments, both strings, and we return two arguments, both strings
func swap(x, y string) (string, string) {
	return y, x
}

// declaring functions with named return values
// i.e. we name what we want to return and then we can just return
// it also establishes these variables as being of that type for you
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// anonymous functions
// note that these have to be declared before they are used
// whereas other functions really are hoisted better
var anonymous = func() {
	fmt.Println("Anonymous function")
}

var anonymous2 = func(x int) (msg string) {
	msg = "Anonymous: " + strconv.Itoa(x)
	fmt.Println(msg)
	return
}

var anonymous3 = func(x int) (msg string) {
	msg = "top secret: "
	for i := 0; i < x; i++ {
		// note that this is better than passing directly,
		// especially if you are dealing with async functions
		func(i int) {
			msg += strconv.Itoa(i)
		}(i)
	}
	return
}

// pointers
func sayName(name *string) {
	fmt.Println(*name) // prints John
	*name = "Peter"
	fmt.Println(*name) // prints Peter
}

// variable number of args (named return)
func sum(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}

// same but with pointers
func pointerSum(nums ...int) *int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return &total
}

func sumArr(nums []*int) (total int) {
	for _, num := range nums {
		total += *num
	}
	return
}

func errors(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("you cannot divide by zero")
	}
	return a / b, nil
}

// passing functions as arguments
func inner() {
	fmt.Println("inner")
}

func outer(function func()) {
	fmt.Println("outer")
	function()
}

// currying functions
func addChain(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// receiver functions (methods)
type String string

func (s *String) tolower() *String {
	*s = String(strings.ToLower(string(*s)))
	return s
}

func (s *String) toupper() *String {
	*s = String(strings.ToUpper(string(*s)))
	return s
}

// function taking a struct as an argument
type person struct {
	name string
	age  int
}

func greet(p person) {
	fmt.Printf("Greetings %v, I am %v years old\n", p.name, p.age)
}

// with a reciever
func (p person) sayName() {
	// note here that you get access to the person struct and its types
	fmt.Printf("Greatings %v, I am %v years old\n", p.name, p.age)
}

// you can also do this with pointers if you are a glutton for punishment

func main() {
	fmt.Println(split(5)) // prints 2 3 (integers)
	//fmt.Println(split(5).x)  // this doesn't work

	anonymous()
	anonymous2(5)
	fmt.Println(anonymous3(5))
	var anonymous4 = func(x int) (msg string) {
		msg = "Anonymous4: " + strconv.Itoa(x)
		fmt.Println(msg)
		return
	}
	anonymous4(5)

	// sepeated declaration and assignment
	var anonymous5 func(int) string
	anonymous5 = func(x int) string {
		return "Anonymous5: " + strconv.Itoa(x)
	}

	fmt.Println(anonymous5(5))

	// pointers
	name := "John"
	sayName(&name)    // pass in the address of the variable
	fmt.Println(name) // prints Peter

	sum(1, 2, 3, 4, 5) // 15
	pointerVal := *pointerSum(1, 2, 3, 4, 5, 6)
	fmt.Println(pointerVal) // 21
	pointers := []*int{&pointerVal, &pointerVal}
	fmt.Println(sumArr(pointers)) // 42

	var nums = []int{1, 2, 3, 4, 5}

	// print the values of nums
	for _, num := range nums {
		fmt.Println(num)
	}

	// get nums as an array of pointers
	var numsPtr []*int
	for i := range nums {
		numsPtr = append(numsPtr, &nums[i])
	}

	fmt.Println(sumArr(numsPtr)) // prints 15
	*numsPtr[0] = 50
	fmt.Println(sumArr(numsPtr)) // prints 64 so this was a little useless really

	fmt.Println(errors(1, 2))
	fmt.Println(errors(1, 0))

	outer(inner) // prints outer inner

	fmt.Println(addChain(5)(10)) // prints 15

	var s String = "HELLO"
	fmt.Println(*s.tolower())           // prints hello
	fmt.Println(*s.tolower().toupper()) // prints HELLO

	p := person{name: "John", age: 30}
	greet(p)
	p.sayName()
}
