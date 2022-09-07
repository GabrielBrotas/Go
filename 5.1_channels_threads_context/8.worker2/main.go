package main

import (
	"fmt"
	"math/rand"
	"time"
)


func GenerateRandomNumber(maxValue int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(maxValue)
}

const poolSize = 10
func CalculateValue(intChan chan int) {
	fmt.Println(4)
	randomNumber := GenerateRandomNumber(poolSize)

	fmt.Println(5)
	intChan <- randomNumber

	fmt.Println(6)
}

func main() {
	intChan := make(chan int)
	fmt.Println(1)
	defer close(intChan) // it will close the channel once it finishes the execution
	// whatever comes after this defer execute as soon as the current function is done
	// so if we open a file or an external connection we dont want to keep that open forever
	
	fmt.Println(2)
	go CalculateValue(intChan)

	fmt.Println(3)
	randomNumber := <- intChan
	fmt.Println(7)

	fmt.Println("Result => ", randomNumber)
}