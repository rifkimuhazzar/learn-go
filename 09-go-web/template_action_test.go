package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func templateActionIf(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/if.html"))
	t.ExecuteTemplate(w, "if.html", map[string]any{
		"Title": "Template action if",
		"Name": "Svelte",
	})
}

func TestTemplateActionIf(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateActionIf(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

func templateActionComparator(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/comparator.html"))
	t.ExecuteTemplate(w, "comparator.html", map[string]any{
		"Title": "Template action comparator",
		"FinalValue": 95,
	})
}

func TestTemplateActionComparator(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateActionComparator(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

func templateActionRange(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/range.html"))
	t.ExecuteTemplate(w, "range.html", map[string]any{
		"Hobbies": []any{
			"Reading",
			"Writing", 
			"Painting", 
			"Photography", 
			nil,
			"Cooking", 
			"Playing Music", 
			"Gardening", 
			"Sports", 
			"Hiking",
		},
		"Title": "Template action range",
	})
}

func TestTemplateActionRange(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateActionRange(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

func templateActionWith(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/with.html"))
	t.ExecuteTemplate(w, "with.html", map[string]any{
		"Title": "Template action with",
		"Name": "JS/TS and GO",
		"Framework": map[string]any{
			"FrontEnd": "Svelte",
			"BackEnd": "Go",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateActionWith(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}