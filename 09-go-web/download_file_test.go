package goweb

import (
	"net/http"
	"testing"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("name")
	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("-- Bad Request --"))
		return
	}
	w.Header().Add("content-disposition", "attachment; filename=\"" + file + "\"")
	http.ServeFile(w, r, "./resources/" + file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(downloadFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}