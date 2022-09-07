package main

import "fmt"

func main() {

	forever := make(chan string)

	go func() {
		/*
		Even tough in this case the scheduler would kill this thread because is taking longer to execute
		the runtime will be stuck in the line 20 waiting for the result var be assigned
		*/
		for {

		}
	}()

	fmt.Println("Awaiting response...")

	result := <- forever // it will be wait forever because result will never be assigned
	
	fmt.Println(result)
}