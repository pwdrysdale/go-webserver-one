package main

import "fmt"

// interfaces are a way to define a set of methods that a type must implement
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

// in go we don't have to explicitly say that a type implements an interface

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}
