package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	link := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//channels create a connection between the main route and sub routes of go
	//the main route of go is when we start application, sub are when we declare them with "go blah blah blah"
	c := make(chan string)

	for _, link := range link {
		go checkLink(link, c)

	}

	//this way we learn that we can send and receive data with channel
	//we can also use a for loop to receive data from the channel
	//with the for loop, we iterate over the length of links, and print the data received from the channel

	// for i := 0; i < len(link); i++ {
	// 	go checkLink(<-c, c)
	// }

	//infinite loop
	for l := range c {
		//go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// you can send and receive data with channel
func checkLink(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
