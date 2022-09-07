package main

import (
	"fmt"
	"time"
)

func main() { // Thread 1
	// a channel will allow the communication between channels

	// T1 <--> T2
	hello := make(chan string) // create a channel that will receive a input as string
	fmt.Println(1)
	
	go func() { // Thread 2
		fmt.Println(3)
		time.Sleep(time.Second*2)
		hello <- "Value from thread 2"
		fmt.Println(4)
	}()
	
	fmt.Println(2)

	// it will stop the thread to await the value
	result := <-hello // when receive a value assign to this variable
	fmt.Println(5)

	fmt.Println(result)
}
