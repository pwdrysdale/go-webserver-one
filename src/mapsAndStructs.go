package main

import (
	"fmt"
	"reflect"
)

// Maps
// Looks like JS / TS maps, but without the set and get methods
// and some of the strange referencing
// The keys for the map must be able to be checked for equality,
// so something like a slice cannot be a key (but arrays can be)

// Structs
// Structs are basically classes or objects in other languages
// They are a collection of fields
// You can create a struct with a type and then create instances of that type
// You don't really need to create duplicates of the struct if you want most
//  of the fields to be the same, you can just create a new instance and
//  override the fields you want to change

func main(){
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 20612439,
		"New York": 19745289,
		"Pennsylvania": 12802503,
		"Illinois": 12801539,
		"Ohio": 11614373,
	} 

	fmt.Println(statePopulations)

	fmt.Println(statePopulations["Florida"])

	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)

	delete(statePopulations, "Georgia")

	//  the behaviour is strange if you try to reference a key that doesn't exist
	// so use the comma ok idiom to check if the key exists
	// if it doesn't exist, the value will be the zero value for the type
	// and the second value will be false
	pop, ok := statePopulations["Georgia"]
	fmt.Println(pop, ok)
	pop2, ok := statePopulations["Florida"]
	fmt.Println(pop2, ok)

	fmt.Println(len(statePopulations))

	//  again you will want to crete a hard duplicate before you start modifying

	var newStatePopulations = make(map[string]int)
	for k, v := range statePopulations {
		newStatePopulations[k] = v
	}

	newStatePopulations["Georgia"] = 5000000
	fmt.Println(statePopulations)
	fmt.Println(newStatePopulations) // Georgia is back, with a new value

	type Doctor struct {
		number int
		actorName string
		companions []string
	}

	aDoctor := Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])
	// fmt.Println(aDoctor["actorName"]) // this doesn't work, you can't use the field name as a key

	// dynamically get a key from a struct
	// you can use the reflect package to get the field names
	// but it's a bit of a pain
	// you can also use the struct tag to get the field name
	// but that's a bit of a pain too
	// so you can just use the field name as a key
	
	// reflect
	t := reflect.TypeOf(aDoctor)
	field, _ := t.FieldByName("actorName") // no success here in getting things dynamically
	// apparently this is an interface

	fmt.Println(field.Name)
	fmt.Println(field.Tag)

	type Animal struct {
		Name string `required max:"100"`
		Origin string
	}

	type Bird struct {
		Animal
		SpeedKPH float32
		CanFly bool
	}

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false

	fmt.Println(b)

	s:= reflect.TypeOf(b)
	fld, _ := s.FieldByName("Name")
	fmt.Println(fld.Tag) // required max:"100" - now can do stuff with its

}