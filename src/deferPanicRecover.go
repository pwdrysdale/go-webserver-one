package main

import "fmt"

func deferExample() {
	// any thing that is marked as defer will execute after
	// everything else in the function has executed
	// but before the function returns
	defer fmt.Println("world")
	fmt.Println("hello")
}

func multiDefer() {
	// defer statements are executed in LIFO order
	defer fmt.Println("world")
	defer fmt.Println("hello")

	// prints hello world
}

func deferOrder() {
	var a = "start"
	defer fmt.Println(a)
	a = "end"

	// prints start => takes the value at the time of defer being called
}

func panicExample() {
	var a, b int = 1, 0
	fmt.Println(a / b)

	// this will panic and stop the program
	// its basically an error message but it
	// stops the program so should be used sparingly
}

func panicDefer() {
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")

	// defer statements are executed before the panic
}

func panicker() {
	fmt.Println("about to panic")

	// recover should only be used in defer functions
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
			panic(err) // if you want
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

func runPanicker() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")

	// prints start, about to panic, error: something bad happened, end
}

func main() {
	deferExample()
	multiDefer()
	deferOrder()
	//	panicExample()
	//panicDefer()
	runPanicker()
}
