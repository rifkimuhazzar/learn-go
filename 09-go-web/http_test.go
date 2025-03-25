package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Method)
	fmt.Fprint(w, "\nHello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest("hello", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()
	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}