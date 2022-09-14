package main

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//func main() { // cannot have more than one main function in a package
//	fmt.Println("What does this do? ")
//}
