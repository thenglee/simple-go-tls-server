package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
