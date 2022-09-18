package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type noun struct {
	name     string
	age      int
	sentient bool
}

type pronoun struct {
	noun
	beingType string
}

var pete = pronoun{
	noun: noun{
		name:     "Pete",
		age:      42,
		sentient: true,
	},
	beingType: "human",
}

func (n noun) describeNoun() string {
	var sentientFlag = ""
	if !n.sentient {
		sentientFlag = "not "
	}

	// if we want the being type
	// get the keys from struct n
	keys := reflect.TypeOf(n)
	// get the values from struct n
	//values := reflect.ValueOf(n)
	fmt.Println(keys)

	return n.name + " is " + strconv.Itoa(n.age) + " years old, and is " + sentientFlag + "sentient."
}

func main() {
	println(pete.name)      // Pete
	println(pete.beingType) // human

	var book = noun{
		name:     "Fear and Trembling",
		age:      42,
		sentient: false,
	}

	fmt.Println(book.describeNoun())
	fmt.Println(pete.describeNoun())
}
