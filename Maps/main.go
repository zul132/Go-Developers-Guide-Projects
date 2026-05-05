package main

import "fmt"

/*

 */

func main() {
	colours := map[string]string{
		"red":   "#ff0000",
		"blue":  "#357ff7",
		"white": "#ffffff",
	}

	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}
