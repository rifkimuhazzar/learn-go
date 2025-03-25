package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func requestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recoder := httptest.NewRecorder()
	requestHeader(recoder, request)

	response := recoder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func responseHeader(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("This-is-custom-header", "Hello World")
	w.Write([]byte("OK"))
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()
	responseHeader(recoder, request)

	response := recoder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	
	fmt.Println(recoder.Header().Get("this-is-custom-header") )
}

