package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func templateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "post.html", map[string]any{
		"Title": "Hello Template Auto Escape",
		"Body": "<p>Hello this is p element</p><script>alert('You got hacked!')</script>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateAutoEscape(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(templateAutoEscape),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func templateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "post.html", map[string]any{
		"Title": "Hello Template Auto Escape",
		"Body": template.HTML("<p>Hello this is p element</p><script>alert('You got hacked!')</script>"),
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateAutoEscapeDisabled(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(templateAutoEscapeDisabled),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func templateXSS(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "post.html", map[string]any{
		"Title": "Template XSS",
		"Body": template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	recorder := httptest.NewRecorder()
	bodyValue := url.QueryEscape("<p>Hello this is p element</p><script>alert('You got hacked!')</script>")
	request := httptest.NewRequest("GET", "localhost:8080?body=" + bodyValue, nil)
	templateXSS(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(templateXSS),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}