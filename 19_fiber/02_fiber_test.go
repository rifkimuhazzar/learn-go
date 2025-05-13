package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
	"github.com/stretchr/testify/assert"
)

var engine = mustache.New("./template", ".mustache")

var app = fiber.New(fiber.Config{
	Views: engine,
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString("Error: " + err.Error())
	},
})

func TestFiber(t *testing.T) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	request := httptest.NewRequest("GET", "/", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(bytes))
}

func TestCtx(t *testing.T) {
	app.Get("/hello", func(ctx *fiber.Ctx) error {
		name := ctx.Query("name", "Guest")
		return ctx.SendString("Hello " + name)
	})

	// Test 1 - with query param
	request := httptest.NewRequest("GET", "/hello?name=Svelte%20Kit", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Svelte Kit", string(bytes))

	// Test 2 - without query param
	request = httptest.NewRequest("GET", "/hello", nil)
	response, err = app.Test(request)
	assert.Nil(t, err)

	bytes, err = io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Guest", string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(ctx *fiber.Ctx) error {
		library := ctx.Get("library")
		metaframework := ctx.Cookies("metaframework")
		return ctx.SendString("Library: " + library + "\nMetaframework: " + metaframework)
	})

	request := httptest.NewRequest("GET", "/request", nil)
	request.Header.Set("library", "Svelte")
	request.AddCookie(&http.Cookie{Name: "metaframework", Value: "Svelte Kit"})

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Library: Svelte\nMetaframework: Svelte Kit", string(bytes))
}

func TestRouteParameter(t *testing.T) {
	app.Get("/users/:userId/order/:orderId", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId")
		orderId := ctx.Params("orderId")
		return ctx.SendString("Get order " + orderId + " from user " + userId)
	})

	request := httptest.NewRequest("GET", "/users/10/order/20", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Get order 20 from user 10", string(bytes))
}

func TestFormRequest(t *testing.T) {
	app.Post("/hello", func(ctx *fiber.Ctx) error {
		name := ctx.FormValue("name")
		return ctx.SendString("Hello " + name)
	})

	requestBody := strings.NewReader("name=Svelte Kit")
	request := httptest.NewRequest("POST", "/hello", requestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Svelte Kit", string(bytes))
}

//go:embed source/contoh.txt
var contohFile []byte

func TestUpload(t *testing.T) {
	app.Post("/upload", func(ctx *fiber.Ctx) error {
		file, err := ctx.FormFile("file")
		if err != nil {
			return err
		}

		err = ctx.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}

		return ctx.SendString("Upload success")
	})

	requestBody := new(bytes.Buffer)
	writer := multipart.NewWriter(requestBody)
	file, err := writer.CreateFormFile("file", "example.txt")
	assert.Nil(t, err)
	file.Write(contohFile)
	writer.Close()

	request := httptest.NewRequest("POST", "/upload", requestBody)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload success", string(bytes))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("/login", func(ctx *fiber.Ctx) error {
		body := ctx.Body()
		request := new(LoginRequest)

		err := json.Unmarshal(body, request)
		if err != nil {
			return nil
		}

		return ctx.SendString("Hello " + request.Username)
	})

	body := strings.NewReader(`{"username": "Svelte", "password": "12345678"}`)
	request := httptest.NewRequest("POST", "/login", body)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Svelte", string(bytes))
}

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"passsword" form:"passsword"`
	Name     string `json:"name" xml:"name" form:"name"`
}

func bodyParser() {
	app.Post("/register", func(ctx *fiber.Ctx) error {
		request := new(RegisterRequest)
		err := ctx.BodyParser(request)
		if err != nil {
			return err
		}
		return ctx.SendString("Register success: " + request.Name)
	})
}

func TestBodyParserJSON(t *testing.T) {
	bodyParser()

	body := strings.NewReader(`{"username": "svelte", "password": "12345678", "name": "Svelte"}`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register success: Svelte", string(bytes))
}

func TestBodyParserForm(t *testing.T) {
	bodyParser()

	body := strings.NewReader("username=vue&password=12345678&name=Vue+Nuxt")
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register success: Vue Nuxt", string(bytes))
}

func TestBodyParserXML(t *testing.T) {
	bodyParser()

	body := strings.NewReader(`
        <RegisterRequest>
            <username>React</username>
            <password>12345678</password>
            <name>React Next</name>
        </RegisterRequest>
    `)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/xml")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register success: React Next", string(bytes))
}

func TestResponseJSON(t *testing.T) {
	app.Get("/user", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"username": "svelte",
			"name":     "Svelte Kit",
		})
	})

	request := httptest.NewRequest("GET", "/user", nil)
	// request.Header.Set("Accept", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"Svelte Kit","username":"svelte"}`, string(bytes))
}

func TestDownloadFile(t *testing.T) {
	app.Get("/download", func(ctx *fiber.Ctx) error {
		return ctx.Download("./source/contoh.txt", "downloaded.txt")
	})

	request := httptest.NewRequest("GET", "/download", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, `attachment; filename="downloaded.txt"`, response.Header.Get("Content-Disposition"))

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Go\n", string(bytes))
}

func TestRoutingGroup(t *testing.T) {
	helloWorld := func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	}

	api := app.Group("/api")
	api.Get("/hello", helloWorld)
	api.Get("/world", helloWorld)

	web := app.Group("/web")
	web.Get("/hello", helloWorld)
	web.Get("/world", helloWorld)

	request := httptest.NewRequest("GET", "/api/hello", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(bytes))
}

func TestStatic(t *testing.T) {
	app.Static("/public", "./source")

	request := httptest.NewRequest("GET", "/public/contoh.txt", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Go\n", string(bytes))
}

func TestErrorHandling(t *testing.T) {
	app.Get("/error", func(ctx *fiber.Ctx) error {
		return errors.New("Terjadi error!")
	})

	request := httptest.NewRequest("GET", "/error", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 500, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Error: Terjadi error!", string(bytes))
}

func TestView(t *testing.T) {
	app.Get("/view", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", map[string]any{
			"title":     "Hello Title",
			"header":    "Hello Header",
			"paragraph": "Hello Paragraph",
		})
	})

	request := httptest.NewRequest("GET", "/view", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "Hello Title")
	assert.Contains(t, string(bytes), "Hello Header")
	assert.Contains(t, string(bytes), "Hello Paragraph")
}

func TestClient(t *testing.T) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	agent := client.Get("https://example.com")
	status, response, errs := agent.String()
	assert.Nil(t, errs)
	assert.Equal(t, 200, status)
	assert.Contains(t, response, "Example Domain")

	fmt.Println(response)
}
