package main

import "fmt"

/*
	The 'Struct' data structure is a collection of properties that are related together.

	Go is a "Pass by Value" language
	--> Anytime we pass a value to a function, either as a receiver or as an argument,
		that data is copied in memory
	--> So that function is always going to be working on a copy of our data structure.
*/

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // alternate way of embedding structs
	// contact   contactInfo
}

func main() {
	/* Note: when declaring structs in multiple lines, ensure that every single line
	ends with a comma ',' even if it is the last property. */

	joe := person{
		firstName: "Joe",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "joe123@gmail.com",
			zipCode: 94000,
		},
	}

	/* joePointer := &joe
	   joePointer.updateName("Joseph") */

	// Pointer shortcut
	joe.updateName("Joseph")

	joe.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// This will NOT update joe's first name
/* func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
} */

func (p person) print() {
	fmt.Printf("%+v", p)
}
