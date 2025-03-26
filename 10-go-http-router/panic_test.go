package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// w.Write([]byte("Hello World"))
		panic("Terjadi Error")
	})
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i any) {
		w.Write([]byte("Panic: " + i.(string)))
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/", nil)
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responseBody))
	assert.Equal(t, "Panic: Terjadi Error", string(responseBody))
}