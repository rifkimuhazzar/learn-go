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

func TestRouterNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Product " + p.ByName("id") + "\n" + "Item " + p.ByName("itemId")
		w.Write([]byte(text))
	})

	recoder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/products/2/items/4", nil)
	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responseBody))
	assert.Equal(t, "Product 2\nItem 4", string(responseBody))
}

func TestCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		imagesPath := p.ByName("image")
		w.Write([]byte("Image: " + imagesPath))
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost:3000/images/png/svelte.png", nil)
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responseBody))
	assert.Equal(t, "Image: /png/svelte.png", string(responseBody))
}