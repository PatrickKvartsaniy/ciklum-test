package main

import (
	"fmt"
	"net/http"

	"ciklum-test/reader/tools"
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
			fmt.Fprintf(w, "Unexpected file, pls choose CSV file")
			return
		}
		defer file.Close()

		StreamCSV(file)
	}
	tools.RenderPage(w, e)
}
