package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func templateLayout(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/content.html",
		"./templates/footer.html",
		"./templates/header.html",
	))
	t.ExecuteTemplate(w, "hello", map[string]any{
		"Title": "Template Layout",
		"Name": "Svelte",
	})
}

func TestTemplateLayout(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateLayout(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}