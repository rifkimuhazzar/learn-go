package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func responseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]uint8("Name is empty"))
	} else {
		// w.WriteHeader(http.StatusOK)
		w.Write([]uint8("Hello " + name))
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	responseCode(recorder, request)

	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}

func TestResponseCodeValid(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080?name=Svelte", nil)
	responseCode(recorder, request)

	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}