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
