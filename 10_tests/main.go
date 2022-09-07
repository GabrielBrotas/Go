package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divide(5, 0)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)

}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Cannot divide by 0")
	}

	result = x / y
	return result, nil
}
