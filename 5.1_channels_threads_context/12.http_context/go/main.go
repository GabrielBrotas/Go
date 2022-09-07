package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func long_execution_query_to_db(ctx context.Context, queryCh chan<- bool) {
	log.Println("Processing...")

	for i := 0; i <= 5; i++ {
		log.Println(i)
		time.Sleep(time.Second)
	}
	queryCh <- true
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		queryCh := make(chan bool)
		go long_execution_query_to_db(ctx, queryCh)
		
		select {
		case result := <-queryCh:
			log.Println("Finish execution")
			response := map[string]interface{}{
				"Success": result,
			}
			reponse_marshelled, _ := json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(reponse_marshelled)
		case <-ctx.Done():
			fmt.Fprint(os.Stderr, "request cancelled\n")
		}

		
	})

	log.Println("App running in port 4000")
	err := http.ListenAndServe(":4000", nil)

	if err != nil {
		log.Fatalln("Error starting the server", err.Error())
	}
}
