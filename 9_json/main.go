package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HasDog    bool `json:"has_dog"`
}

func main() {

	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"has_dog": false
		}
	]
	`

	// READ ------------
	var unmarshalled []Person

	// unmarshal takes a slice of bytes and the interface this the data will be unmarshalled
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("Unmarshalled: %v", unmarshalled)

	// WRITE ------------
	var mySlice []Person

	tony := Person{
		FirstName: "Tony",
		LastName: "Stark",
		HasDog: false,
	}

	peter := Person{
		FirstName: "Peter",
		LastName: "Parker",
		HasDog: false,
	}

	mySlice = append(mySlice, tony, peter)

	newJson, err := json.Marshal(&mySlice)

	if err != nil {
		log.Println("Error marshalling json", err)
	}

	log.Println("Marshalling", string(newJson))

}
