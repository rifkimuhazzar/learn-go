package goweb_test

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func simpleTemplate(w http.ResponseWriter, r *http.Request) {
	templateText := "<html><body>{{.}}</body></html>"
	// t, err := template.New("Simple").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }
	t := template.Must(template.New("Simple").Parse(templateText))
	t.ExecuteTemplate(w, "Simple", "Hello this is simple template!")
}

func TestSimpleTemplate(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	simpleTemplate(recorder, request)

	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))
}

func TestSimpleTemplate2(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(simpleTemplate),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func simpleTemplateFile(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/simple.html"))
	t.ExecuteTemplate(w, "simple.html", "Hello this is file template!")
}

func TestSimpleTemplateFile(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	simpleTemplateFile(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

func templateDirectory(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseGlob("templates/*.html"))
	t.ExecuteTemplate(w, "simple.html", "Hello this is file template!")
}

func TestTemplateDirectory(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateDirectory(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

//go:embed templates/*.html
var templates embed.FS

func templateEmbed(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.html"))
	t.ExecuteTemplate(w, "simple.html", "Hello this is template embed!")
}

func TestTemplateEmbed(t *testing.T) {
	recoder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateEmbed(recoder, request)

	responseBody, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}