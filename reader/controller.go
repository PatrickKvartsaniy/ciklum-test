package main

import (
	"fmt"
	"net/http"
)

// Reader is "/ " handler
func Reader(w http.ResponseWriter, req *http.Request) {
	var e string
	if req.Method == http.MethodPost {
		file, header, err := req.FormFile("uploadfile")
		fileType := header.Header["Content-Type"][0]
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if fileType != "text/csv" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Unexpected file, pls choose CSV file")
			return
		}
		defer file.Close()

		result, err := StreamCSV(file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Something went wrong :/. Error: ", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, result)
		return
	}
	RenderPage(w, e)
}
