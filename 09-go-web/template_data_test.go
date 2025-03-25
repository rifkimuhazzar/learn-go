package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func templateDataMap(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(w, "name.html", map[string]any{
		"Title": "Template Data Map",
		"Name": "Svelte",
		"Address": map[string]any {
			"Street": "USA",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateDataMap(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

type Page struct {
	Title, Name string
	Address
}

type Address struct{
	Street string
}

func templateDataStruct(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(w, "name.html", Page{
		Title: "Template Data Struct",
		Name: "Svelte",
		Address: Address{
			Street: "USA",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateDataStruct(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}