package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "My-Cookie"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"
	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success create cookie")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("My-Cookie")
	if err != nil {
		fmt.Fprint(w, "There is no Cookie")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", setCookie)
	mux.HandleFunc("/get-cookie", getCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080?name=Svelte", nil)
	setCookie(recorder, request)

	responseCookies := recorder.Result().Cookies()
	for _, cookie := range responseCookies {
		fmt.Printf("Cookie - %s: %s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "My-Cookie"
	cookie.Value = "Svelte"
	request.AddCookie(cookie)
	getCookie(recorder, request)

	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))
}