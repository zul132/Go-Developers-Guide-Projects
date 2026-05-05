package main

import "fmt"

/*
	The 'Struct' data structure is a collection of properties that are related together.

	Go is a "Pass by Value" language
	--> Anytime we pass a value to a function, either as a receiver or as an argument,
		that data is copied in memory
	--> So that function is always going to be working on a copy of our data structure.


	VALUE TYPES:
	int, string, float, bool, structs --> Use pointers to change these things in a function

	REFERENCE TYPES:
	slices, maps, channels, pointers, functions --> Don't worry about pointers with these

	When you create a new reference type like a slice, Go will actually create two data structures in memory
	- an array which contains the list of elements,
	- and the slice which contains a pointer to the head of that array.

	When you pass a slice or similar reference type to a function, then Go will still create a copy of your slice in memory
	BUT that copy is going to reference or point to the exact same array (memory address) as your original slice.

	Hence, when you pass a slice or a reference type value to a function and make modifications to that slice inside the function,
	the changes you made are reflected in your original slice even outside the function.
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
