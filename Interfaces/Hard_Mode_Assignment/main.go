package main

import (
	"fmt"
	"io"
	"os"
)

/*
	Task:
	- Read the contents of a text file and print its contents to the terminal.
	- The file to open should be provided as an argument to the program when it is executed at the terminal.

	Solution:
	To read a list of command-line arguments provided to a program, you can reference
	the variable 'os.Args', which is a slice of type string.

	To open a file, we can use the os.Open() function.

	Since the 'File' type implements the 'Reader' interface, we can reuse the io.Copy() function.
*/

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		// log.Fatal(err)
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
