package goweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.html
var templates2 embed.FS

var myTemplate = template.Must(template.ParseFS(templates2, "templates/*.html"))

func templateCaching(w http.ResponseWriter, _ *http.Request) {
	myTemplate.ExecuteTemplate(w, "simple.html", "Hello Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateCaching(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}