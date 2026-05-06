package main

import "fmt"

/*
	Differences between Maps and Structs:-

	1. Map - All keys and values must be of the same type
	   Struct - Values can be of different type

	2. Map - Keys are indexed, we can iterate over them
	   Struct - Keys don't support indexing (we cannot iterate over the values of a struct)

	3. Map - Reference type
	   Struct - Value type

	4. Map - Used to represent a collection of related properties
	   Struct - Struct is used to represent a "thing" (object) with a lot of different properties

	5. Map - You don't need to know all the keys at compile time
	   Struct - You need to know all the different fields at compile time
*/

func main() {
	colours := map[string]string{
		"red":   "#ff0000",
		"blue":  "#357ff7",
		"white": "#ffffff",
	}

	printMap(colours)

	fmt.Println(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}
