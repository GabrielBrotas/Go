package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// channel is the only way that can communicate between go routines and main routine
	// channel is a communication mechanism that allows two or more goroutines to communicate with each other and synchronize their execution
	// create a channel
	c := make(chan string) // create a channel that will receive a string as input

	for _, url := range links {
		// checkLink(url) // this function will generate a block io, will await for the result and then check the next link

		// create a go routine to check the current link and then continue with the next link
		// create a new thread go routine to run the function
		// because go keywork create child routine, it will not block the main routine and will not print in the main thread terminal the result
		go checkLink(url, c)
	}

	// fmt.Println(<- c) // will block the main thread and wait for the channel send a message as a result
	

	// for i := 0; i < len(links); i++ {
	// for { // this is a infinite loop
	//	// fmt.Println(<-c)
		
	// 	// when we get the response link from the channel we will execute again to create a loop and keep checking the link
	// 	go checkLink(<-c, c)
	// }

	// is the same as the code above, 
	// wait for the channel return a value and assign to the l variable
	for l := range c {
		// we should not put sleep function in our main routine because we will no be able to access the channel response 
		// go checkLink(l, c)

		// anonymous function
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}
	
}

/*
	sending data with channels
	channel <- 5 = send the value '5' to this channel
	myNumber <- channel = wait for a value to be sent into the channel, when we get one, assign the value to "myNumber"
	fmt.Println(<- channel) = wait for a value to be sent into the channel, when we get one, print it
*/

func checkLink(link string, channel chan string) {
	// stop the routine in 5 seconds
	// time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	
	if err != nil {
		fmt.Println(link, "might be down")

		// send this message to the channel
		// channel <- "might be down"
		channel <- link
		return
	}

	fmt.Println(link, "is up!")

	// send message to the channel
	// channel <- "is up!"
	channel <- link

}