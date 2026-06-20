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
	and only then do we move on to the next link. --> Can be an issue if we have many links

	- Since requests are made in serial, the links get printed in the same order as they appear in the slice.

	Concurrent Approach:
	- Rather than running the requests in serial, we run every request in parallel and whichever request
	returns a response, we immediately log it.

	- Links get printed in the order in which we get the response, which may differ from the slice order.
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
		checkLink(link)
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
