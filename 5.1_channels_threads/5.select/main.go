package main

import (
	"fmt"
	"time"
)

func main() {

	hello := make(chan string)

	go func() {
		time.Sleep(time.Second)
		hello <- "From Thread 2"
	}()

	select {
	case x := <-hello:
		fmt.Println(x)
	// if we have this default our case above will never be executed
	// because the select is just executed once
	// default:
	// 	fmt.Println("defult")
	}

	time.Sleep(time.Second*5)
}
