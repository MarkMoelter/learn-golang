package main

import (
	"log"
	"net/http"

	"example.com/myproject/greet"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(greet.MyGreeterHandler)))

}
