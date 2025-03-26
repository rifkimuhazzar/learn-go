package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFiles(t *testing.T) {
	router := httprouter.New()
	directory, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}
	router.ServeFiles("/files/*filepath", http.FS(directory))

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responseBody))
	assert.Equal(t, "Hello HTTP Router", string(responseBody))
}

func TestServeFiles2(t *testing.T) {
	router := httprouter.New()
	directory, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}
	router.ServeFiles("/files/*filepath", http.FS(directory))

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost:3000/files/goodbye.txt", nil)
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responseBody))
	assert.Equal(t, "Good Bye HTTP Router", string(responseBody))
}