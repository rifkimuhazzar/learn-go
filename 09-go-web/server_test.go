package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}
	
	err := server.ListenAndServe()
	fmt.Println("Hello World")
	if err != nil {
		panic(err)
	}
}