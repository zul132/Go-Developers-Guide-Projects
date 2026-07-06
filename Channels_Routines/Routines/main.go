package main

import (
	"fmt"
	"net/http"
)

/*
	Website Status Checker:
	Channels and Routines are both Go structures that are used for handling concurrent programming.

	Serial or Sequential Approach (without concurrency):
	- For every link, we make a request and then wait till we get a response before logging the response
	and only then do we move on to the next link.
	- If we had many URLs in our program, then it'll take us a very long time to fetch all the URLs
	  as the requests are made one-by-one.

	- Since requests are made in serial, the links get printed in the same order as they appear in the slice.

	Concurrent Approach:
	- Rather than running the requests in serial, we run every request in parallel and whichever request
	returns a response, we immediately log it.

	- Links get printed in the order in which we get the response, which may differ from the slice order.

	Go Routines:
	We can take any function call that we want to be executed inside of its own individual Go Routine
	and just add the 'go' keyword in front of it.

	Using 'go' creates a new thread Go Routine and runs the function with it.
*/
/*
	THE BUG:
	When we first start up our program, we get this single Main go routine. Anytime we use the 'go' keyword,
	we are creating Child routines.

	Main routine is the single routine which controls when the program quits or exits.

	After finishing with the for loop, the Main routine thinks there is no more code to be run and exits the program.
	It doesn't care that the Child routines haven't finished fetching their URLs yet.

	Hence why nothing gets printed when we run the program.
*/

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		// Use the 'go' keyword to run the checkLink function inside a brand new Go routine
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
