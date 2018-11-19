package main

import (
	"log"
	"net/http"
)

const (
	port = "8080"
)

func main() {

	http.HandleFunc("/", Reader)
	log.Println("Server running on 127.0.0.1:" + port)
	http.ListenAndServe(":"+port, nil)

}
