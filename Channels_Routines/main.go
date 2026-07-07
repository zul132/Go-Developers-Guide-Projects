package main

import (
	"fmt"
	"net/http"
)

/*
	Channels:
	Channels are a construct in Go used to communicate between different running go routines.

	Channels are a typed construct - the data or messages that we send thru a channel must always be of
	the same type.
*/

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// string is the type of data we want to communicate over this channel
	c := make(chan string)

	for _, link := range links {
		// Use the 'go' keyword to run the checkLink function inside a brand new Go routine
		go checkLink(link, c)

		/* If we had printed the value from the channel here instead of outside the for loop
		   then we would've seen all the links printed out in order.
		*/
		// fmt.Println(<-c)
	}

	/* Whenever we wait for a message to come thru the channel --> the Main routine pauses execution and waits.
	   i.e. receiving messages from a channel is a Blocking call

	   The reason why we see only 1 link being printed is - Whichever child routine fetches the URL first will pass
	   the link into the channel --> Main routine receives this link and prints it --> No more lines of code to run
	   --> Main exits the program
	*/
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down i think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep its up"
}
