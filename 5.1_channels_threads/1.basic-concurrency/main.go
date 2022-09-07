package main

import (
	"fmt"
	"time"
)

func counter(prefix string) {
	for i := 0; i <= 10; i++ {
		fmt.Println(prefix, i)
	}
}

func concurrency_counter(prefix string) {
	for i := 0; i <= 10; i++ {
		fmt.Println(prefix, i)
		time.Sleep(time.Second)
	}
}

func main() { // Thread 1

	fmt.Println("Hello 1")
	counter("sync")
	go counter("another thread") // Thread 2

	time.Sleep(time.Second * 2) // we need await the second thread be executed to see the logs
	fmt.Println("End----------")

	// Concurrency execution
	go concurrency_counter("thread 2") 
	go concurrency_counter("thread 3")
	time.Sleep(time.Second * 10)
}
