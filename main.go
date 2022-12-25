package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("received request")
	fmt.Fprintf(w, "hello world!")
}

func main() {
	http.HandleFunc("/", healthCheck)

	log.Println("Start server")

	server := &http.Server{
		Addr: ":8080",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
