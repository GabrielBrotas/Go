package main

import (
	"fmt"
	"sync"
	"time"
)

func generateNumbers(total int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= total; i++ {
		fmt.Printf("Generating number %d\n", i)
	}
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // defer to call Done to decrease the count by one after the function finishes running.

	for i := 1; i <= 3; i++ {
		fmt.Printf("Printing number %d\n", i)
		time.Sleep(time.Second)
	}
}


/*
To wait for the functions to finish, you’ll use a WaitGroup from Go’s sync package. 
The sync package contains “synchronization primitives”, such as WaitGroup, that are designed to synchronize various parts of a program. 
In your case, the synchronization keeps track of when both functions have finished running so you can exit the program
*/
func main() {
	var wg sync.WaitGroup // The WaitGroup primitive works by counting how many things it needs to wait for using the Add, Done, and Wait functions.
	
	// The Add function increases the count by the number provided to the function, and Done decreases the count by one. 
	wg.Add(2)
	go printNumbers(&wg)
	go generateNumbers(3, &wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait() // The Wait function can then be used to wait until the count reaches zero, meaning that Done has been called enough times to offset the calls to Add. Once the count reaches zero, the Wait function will return and the program will continue running.
	// After declaring the WaitGroup, it will need to know how many things to wait for. Including a wg.Add(2) in the main function before starting the goroutines will tell wg to wait for two Done calls before considering the group finished
	
	fmt.Println("Done!")
}