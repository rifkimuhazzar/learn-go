package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func formPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	// r.PostFormValue()
	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")
	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	recoder := httptest.NewRecorder()
	requestBody := strings.NewReader("first_name=React&last_name=Vue")
	request := httptest.NewRequest("POST", "http://localhost:8080", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	formPost(recoder, request)

	response := recoder.Result()
	responseBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))

	fmt.Println(response.Header.Get("content-type")) 
}

