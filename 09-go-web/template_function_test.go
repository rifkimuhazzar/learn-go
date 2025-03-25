package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (m MyPage) SayHello(name string) string {
	return "Hello " + name + ", my name is " + m.Name
}

func templateFunction(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.New("function").Parse(`<p>{{.SayHello "React"}}</p>`))

	err := t.ExecuteTemplate(w, "function", MyPage{"Svelte"})
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunction(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateFunction(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

func templateFunctionGlobal(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.New("function").Parse(`<p>{{len .Name}}</p>`))

	err := t.ExecuteTemplate(w, "function", MyPage{"Svelte"})
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunctionGlobal(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateFunctionGlobal(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

func templateFunctionGlobalCustom(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(
		template.New("function").
		Funcs(map[string]any{
			"upper": func(value string) string  {
				return strings.ToUpper(value)
			},
		}).
		Parse(`<p>{{upper .Name}}</p>`),
	)

	err := t.ExecuteTemplate(w, "function", MyPage{"Svelte"})
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunctionGlobalCustom(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateFunctionGlobalCustom(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

func templateFunctionPipelines(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(
		template.New("function").
		Funcs(map[string]any{
			"upper": func(value string) string  {
				return strings.ToUpper(value)
			},
			"sayHello": func(value string) string  {
				return "Hello " + value
			},
		}).
		Parse(`<p>{{sayHello .Name | upper}}</p>`),
	)
	
	err := t.ExecuteTemplate(w, "function", MyPage{"Svelte"})
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunctionPipelines(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	templateFunctionPipelines(recorder, request)

	responseBody, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}