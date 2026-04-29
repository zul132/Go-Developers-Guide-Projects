package main

import "fmt"

/*
	The 'Struct' data structure is a collection of properties that are related together.
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

	joe.updateName("Joseph")
	joe.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
