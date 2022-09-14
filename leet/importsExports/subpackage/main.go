package subpackage

import "fmt"

var Sub string = "Sub"

func SomeFunction() {
	fmt.Println("someFunction is printing this")
}

func Subpackage() {
	fmt.Println("Hello from subpackage")
}
