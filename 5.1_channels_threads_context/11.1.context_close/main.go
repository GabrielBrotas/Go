package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func write_generateNumbers(ctx context.Context,  wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancelCtx := context.WithCancel(ctx) // provides a function to close the context, it's the same thing as if the client close a http request, the context will close

	numberChan := make(chan int)
	
	go reader_printNumbers(ctx, numberChan, wg)

	for num := 1; num <= 5; num++ {
		numberChan <- num
		time.Sleep(time.Second)

	}

	cancelCtx() // it will break the reader infinity loop

	fmt.Printf("Writer: finished Sucessfully\n")
}

func reader_printNumbers(ctx context.Context, numberChan <-chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()

			if err != nil {
				fmt.Printf("Reader err: %s\n", err)
			} else {
				fmt.Printf("Reader: finished successfully\n")
			}
			defer wg.Done()
			return
		case num := <-numberChan:
			fmt.Printf("Reader: %d\n", num)
		}
	}
}

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup
	
	wg.Add(2)

	log.Println("Running program...")
	write_generateNumbers(ctx, &wg)

	wg.Wait()

}
