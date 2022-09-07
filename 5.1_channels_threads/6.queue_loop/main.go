package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int)
	
	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			/*
			First it will assign the number 0 but then the code will be stuck
			until a consumer read this value, otherwise this loop will never be
			executed again
			*/
			queue <- i // each time this queue is read by a consumer another interaction will occur
			i++
		}
	}()

	for x := range queue {
		fmt.Println(x)
	}

}