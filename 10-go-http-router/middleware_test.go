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

type logMiddleware struct{
	http.Handler
}

func (lm *logMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive request!")
	lm.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("Hello World"))
	})

	middleware := logMiddleware {
		Handler: router,
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/", nil)
	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responseBody))
	assert.Equal(t, "Hello World", string(responseBody))
}