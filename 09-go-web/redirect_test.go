package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func redirectFrom(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func redirectOut(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://youtube.com", http.StatusTemporaryRedirect)
}

func redirectTo(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello Redirect")
}

func TestRedirect(t *testing.T) {
	mux := new(http.ServeMux)
	mux.HandleFunc("/redirect-from", redirectFrom)
	mux.HandleFunc("/redirect-out", redirectOut)
	mux.HandleFunc("/redirect-to", redirectTo)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}