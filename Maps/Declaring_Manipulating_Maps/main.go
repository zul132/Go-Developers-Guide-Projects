package main

import "fmt"

/*
	Maps in Go are a collection of key-value pairs.

	Both the keys and values are statically typed
	- All the keys are of the same type
	- All the values are of the same type

	Maps are an Unordered data structure.

	There are a few different ways of declaring and manipulating a Map in Go.
*/

func main() {
	// method 1
	/* colours := map[string]string{
		"red":   "#ff0000",
		"blue":  "#357ff7",
		"green": "#4bf745",
	} */

	// method 2: using the 'var' keyword
	//var colours map[string]string

	// method 3: Using the built-in go function 'make'
	colours := make(map[int]string)

	colours[10] = "#ffffff"

	delete(colours, 10) // 'delete' is a built-in go function

	fmt.Println(colours)
}
