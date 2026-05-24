package main

import (
	"fmt"
	"net/http"
	"os"
)

/*
	More Interface Syntax:
	- In Go, we can take multiple interfaces and assemble them together to form another interface.
	- We use interfaces to reuse common code.

	Reader interface provides a common output []byte for all the different sources of input data.

	Sources of Input:
	HTTP Request Body,
	Text file on hard drive,
	Image file on hard drive,
	User entering text into command line
	Data from analog sensor plugged into machine

	type Reader interface {
		Read(p []byte) (n int, err error)
		|
	    --> Read() accepts a byte slice and pushes data into the byte slice
	}

	Whoever wants to consume data from a source creates their own byte slice, which they pass into Read()
	and then Read() will push some data into the byte slice.

	n -> number of bytes read into the slice

	Difference between other languages & Go:
	- You'd expect Read() to directly return a byte slice in other languages.
	- But in Go, Read() accepts a byte slice from the caller and pushes data into the byte slice.
*/

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// We create a reasonably gigantic sized byte slice of zeros using make() and pass it into Read()
	// Reason: If we simply passed in an empty slice, Read() will say that
	// the slice is already full and won't ready any data into it.
	bs := make([]byte, 99999)
	resp.Body.Read(bs)

	fmt.Println(string(bs))
}
