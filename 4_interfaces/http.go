package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// give me splice that contains 99999 bytes of empty string 0-n
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// ------

	// this io.Copy implements the Writer and Reader interfaces
	// func Copy(dst Writer, src Reader) (written int64, err error)
	// io.Copy(os.Stdout, resp.Body)

	//  --------------------------------
	logW := logWriter{}

	io.Copy(logW, resp.Body)
}


func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}