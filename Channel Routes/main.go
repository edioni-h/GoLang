package main

import (
	"fmt"
	"net/http"
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
	fmt.Println(<-c)
}

// you can send and receive data with channel
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be down!")
		c <- "Might be down..."
		return
	}

	fmt.Println(link, "is up!")
	c <- "It is up and running!"
}
