package main

import (
	"fmt"
	"net/http"
)

func Reader(w http.ResponseWriter, req *http.Request) {
	// if we got a file
	if req.Method == http.MethodPost {
		// start file uploading
		file, header, err := req.FormFile("uploadfile")
		defer file.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// if we got an unexpected file extension
		if header.Header["Content-Type"][0] != "text/csv" {
			// return BadRequest
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Unexpected file, pls choose CSV file")
			return
		}
		// Start streaming
		result, err := StreamCSV(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Fprintf(w, "Something went wrong. Message from server:", result)
			return
		}
		// if everything is fine
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, result)
		return
	} else if req.Method == http.MethodGet {
		RenderPage(w)
	} else {
		http.Error(w, "405 Status Not Allowed", http.StatusMethodNotAllowed)
	}
}
