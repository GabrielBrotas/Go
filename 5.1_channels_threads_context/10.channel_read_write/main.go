package main

import (
	"fmt"
	"sync"
	"time"
)

/*
The position of the <- operator in relation to the channel variable determines whether you’re reading from or writing to the channel.

To write to a channel, begin with the channel variable, followed by the <- operator, then the value you want to write to the channel:
intChan := make(chan int)
intChan <- 10

func writeChannel(ch chan<- int) {
	ch is write-only
}

To read a value from a channel, begin with the variable you want to put the value into, either = or := to assign a value to the variable, followed by the <- operator,
and then the channel you want to read from:
intChan := make(chan int)
intVar := <- intChan

func readChannel(ch <-chan int) {
	ch is read-only
}

If the function declaration doesn’t have an arrow, as in the case of chan int, the channel can be used for both reading and writing.

Finally, once a channel is no longer being used it can be closed using the built-in close() function.
This step is essential because when channels are created and then left unused many times in a program, it can lead to what’s known as a memory leak.
A memory leak is when a program creates something that uses up memory on a computer, but doesn’t release that memory back to the computer once it’s done using it.

 This leads to the program slowly (or sometimes not so slowly) using up more memory over time, like a water leak. When a channel is created with make(),
 some of the computer’s memory is used up for the channel, then when close() is called on the channel, that memory is given back to the computer to be used for something else.

*/

func writer_generateNumbers(total int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= total; i++ {
		fmt.Printf("sending %d to channel\n", i)
		ch <- i
		time.Sleep(time.Second)
	}
}

func reader_printNumbers(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("read %d from channel\n", num)
	}
}
func main() {
	var wg sync.WaitGroup
	numberChan := make(chan int)

	wg.Add(2)
	go reader_printNumbers(numberChan, &wg)

	writer_generateNumbers(3, numberChan, &wg)

	// close() causes the for ... range loop in printNumbers to exit. Since using range to read from a channel continues until the channel it’s reading from is closed, if close isn’t called on numberChan then printNumbers will never finish
	close(numberChan) // close because is no longer being used

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
	
}