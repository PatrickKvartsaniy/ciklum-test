package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	port := flag.String("port", "8080", "port to run")
	flag.Parse()

	http.HandleFunc("/", Reader)
	log.Println("Server running on 127.0.0.1:" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}
