package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() { // Thread 1
	runtime.GOMAXPROCS(1) // force the number of cpus to 1 so we can test Go concurrency

	fmt.Println("Begin...")

	go func() {
		/*
			In go version <= 13 this function would block the main thread
			because go works with cooperative scheduler and this is a infinity loop
			so it will never release the CPU to the main thread execute again

			but since version >= 14, go added some resources that when a situation
			like this happen it will work in a preemptive way
			so if the thread is taking longer to execute the go scheduler will force to switch the context
		*/
		for {

		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Finish...")
}