package main

import "fmt"

/*
	Problems without Interfaces:
	In this program, we will see a simple example of problems we may face without using interfaces.

	We use interfaces to define a method set (set of functions and return types) that a type should have.
	Any type that matches the description of the interface is considered as
	an honorary member of that interface, or a type of that interface.

	So in our example, 'englishBot' and'spanishBot' are both treated as a type of 'bot'.
	This means that we can use these types at any location where a 'bot' would be expected.
*/

/*
	There are two different kinds of "types":-

	1. Concrete Types
	   We can directly create values out of a concrete type, which we can then access and modify, create copies of etc.
	   Eg. map, struct, int, string or custom types like englishBot

	2. Interface Type
	   We can't directly create values out of an interface type.
	   Eg. bot

	Notes:
	- Interfaces are NOT 'generic' types (Go does not have 'generic' types unlike other languages).
	- Interfaces are 'implicit'.
	  i.e. We don't manually have to declare that our custom type satisfies some interface.
	- Interfaces are a contract to help us manage types.
	  GARBAGE IN -> GARBAGE OUT. They cannot help prevent incorrect implementation of a function's logic.
*/

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// We can omit the "eb" value from (eb englishBot) in our function definition
// as we are not actually going to use "eb" inside our function logic
func (englishBot) getGreeting() string {
	// VERY custom logic for generating an English greeting
	return "Hi there!"
}

// Similarly, (spanishBot) is equivalent to writing (sb spanishBot)
func (spanishBot) getGreeting() string {
	// VERY custom logic for generating a Spanish greeting
	return "Hola!"
}

// Go does NOT support method overloading
// i.e. We cannot declare functions with identical names in the same file
/*
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
*/
