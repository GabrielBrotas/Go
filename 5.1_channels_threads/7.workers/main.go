package main

import (
	"fmt"
	"time"
)


func worker(workerId int, msg chan int) {
	// infinity loop through the received messages
	for res := range msg {
		fmt.Println("Worker: ", workerId, "Msg: ", res)
		time.Sleep(time.Second)
	}
}

func main() {
	msg := make(chan int)

	go worker(1, msg)
	go worker(2, msg)

	for i:=0; i<=10;i++ {
		// we need to wait the worker to respond before send another message
		// to increase the throughput we can add more workers
		msg <- i // it will send the message to the worker
		
	}
}