package main

import "fmt"

// constants are created at compoile time, hence they cannot
// be created using functions or anything that needs to be run at runtime

// they are replaced by the compiler at compile time

const (
	a = 42
	b = 42.78
)

// iota is an integer which is incremented by 1 each time it is used
// and constants follow patterns of the block

const (
	c = iota
	d
	e
)

// This can be handy when using the bitshift to essentially create enums
const (
	admin =  1 << iota
	user
	other

	canSeeAustralia
	canSeeAfrica
	canSeeAsia
	canSeeEurope
)

func main() {
	fmt.Println(a)
	const b = 43
	fmt.Println(b) // local scope overwrites the global scope

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println(admin)
	fmt.Println(user)
	fmt.Println(other)

	var roles byte = admin | canSeeAustralia | canSeeAsia

	fmt.Printf("%b\n", roles) // 101001 essentially can sse Asia, Australia and is an admin => the powers of two for the roles

	fmt.Printf("Is admin? %v\n", admin & roles == admin)
	fmt.Printf("Is user? %v\n", user & roles == user)
}