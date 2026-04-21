package main

import "fmt"

/*
	The 'Struct' data structure is a collection of properties that are related together.
*/

type person struct {
	firstName string
	lastName  string
}

func main() {
	/*
		There are 3 ways of declaring Struct values
	*/

	// Method 1
	// alex := person{"Alex", "Anderson"}

	// Method 2
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var alexa person

	alexa.firstName = "Alexa"
	alexa.lastName = "Anderson"

	fmt.Println(alexa)
	fmt.Printf("%+v", alexa)
}
