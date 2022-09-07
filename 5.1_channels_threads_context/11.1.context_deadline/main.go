package main

import (
	"context"
	"log"
	"sync"
	"time"
)

/*
	Using context.WithDeadline with a context allows you to set a deadline for when the context needs to be finished,
	and it will automatically end when that deadline passes.
	Setting a deadline for a context is similar to setting a deadline for yourself.
	You tell the context the time it needs to be finished, and Go automatically cancels the context for you if that time is exceeded.
*/

func writer_generateNumbers(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	deadline := time.Now().Add(time.Second * 5) // 5 seconds to execute

	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	defer cancelCtx()

	numberChan := make(chan int)

	go reader_printNumbers(ctx, numberChan, wg)

	// it will break after 5 seconds because te dealine will be exceeded
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
