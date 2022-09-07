package main

import (
	"context"
	"log"
	"sync"
	"time"
)

/*
	The context.WithTimeout function can almost be considered more of a helpful function around context.WithDeadline.
	With context.WithDeadline you provide a specific time.Time for the context to end, but by using the context.WithTimeout function you only need to provide
	a time.Duration value for how long you want the context to last. This will be what you’re looking for in many cases,
	but context.WithDeadline is available if you need to specify a time.Time.
	Without context.WithTimeout you would need to use time.Now() and the time.Time’s Add method on your own to get the specific time to end.
*/

func writer_generateNumbers(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancelCtx := context.WithTimeout(ctx, time.Second*5)
	defer cancelCtx()

	numberChan := make(chan int)

	go reader_printNumbers(ctx, numberChan, wg)

	// it will break after 5 seconds because the timeout will be exceeded
	for num := 1; num <= 10; num++ {
		err := ctx.Err()

		if err != nil {
			break
		}

		select {
		case numberChan <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			log.Println("Writer: Ctx done")
		}
	}

	cancelCtx()

	log.Printf("db_query_that_takes_long_to_execute: finished\n")

}

func reader_printNumbers(ctx context.Context, numberChan <-chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()

			if err != nil {
				log.Printf("Reader err: %s\n", err)
			} else {
				log.Printf("Reader: finished successfully\n")
			}
			wg.Done()
			return
		case num := <-numberChan:
			log.Printf("Reader: %d\n", num)
		}
	}
}
func main() {
	ctx := context.Background()
	var wg sync.WaitGroup

	wg.Add(2)

	log.Println("Running program...")
	writer_generateNumbers(ctx, &wg)

	wg.Wait()
}
