package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTP(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(Reader)

	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("Server didn't respond. HTTP code %v", res.Code)
	}

}

// func TestFileLoading(t *testing.T){

// }
