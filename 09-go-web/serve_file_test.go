package goweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func serveFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") == "" {
		http.ServeFile(w, r, "resources/not-found.html")
	} else {
		http.ServeFile(w, r, "./resources/ok.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(serveFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/not-found.html
var notFound string

//go:embed resources/ok.html
var ok string

func serveFileEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") == "" {
		fmt.Fprint(w, notFound)
	} else {
		// fmt.Fprint(w, ok)
		w.Write([]byte(ok))
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(serveFileEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}