package goweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func uploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplate.ExecuteTemplate(w, "upload.form.html", nil)
	if err != nil {
		panic(err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(32 << 20)
	file, fileHeader, err := r.FormFile("input-file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	defer fileDestination.Close()
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")

	myTemplate.ExecuteTemplate(w, "upload.success.html", map[string]any{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", uploadForm)
	mux.HandleFunc("/upload", upload)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/svelte-vertical.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	requestBody := new(bytes.Buffer)

	writer := multipart.NewWriter(requestBody)
	writer.WriteField("name", "Svelte")
	file, _ := writer.CreateFormFile("input-file", "UploadFileTest.png")

	// with embed
	file.Write(uploadFileTest)

	// without embed
	// fileBytes, _ := os.ReadFile("./resources/svelte-vertical.png")
	// file.Write(fileBytes)

	writer.Close()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "localhost:8080", requestBody)
	request.Header.Set("content-type", writer.FormDataContentType())
	upload(recorder, request)

	responseBody := recorder.Result().Body
	defer responseBody.Close()

	result, err := io.ReadAll(responseBody)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}