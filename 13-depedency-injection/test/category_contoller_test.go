package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"go_restful_api/app"
	"go_restful_api/controller"
	"go_restful_api/helper"
	"go_restful_api/middleware"
	"go_restful_api/models/domain"
	"go_restful_api/repository"
	"go_restful_api/service"
)


func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:n3txt.vml1@tcp(localhost:3306)/learn_go_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	categoryRepository := repository.NewCategoryRepository()
	validate := validator.New()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	_, err := db.Exec("TRUNCATE categories")
	if err != nil {
		panic(err)
	}
}

func TestCreateCategory_WithValidInput_ShouldSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	requestBody := strings.NewReader(`{"name": "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "SECRET")

	router.ServeHTTP(recorder, request)
	
	response := recorder.Result()
	defer response.Body.Close()
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	responseBody := map[string]any{}
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]any)["name"])
}

func TestCreateCategory_WithInvalidInput_ShouldFail(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "SECRET")

	router.ServeHTTP(recorder, request)
	
	response := recorder.Result()
	defer response.Body.Close()
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	responseBody := map[string]any{}
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetByIdCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepository()
	tx, _ := db.Begin()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, 
	"http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-KEY", "SECRET")

	router.ServeHTTP(recorder, request)
	
	response := recorder.Result()
	defer response.Body.Close()
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	responseBody := map[string]any{}
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]any)["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]any)["name"])
}

func TestGetByIdCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("X-API-KEY", "SECRET")

	router.ServeHTTP(recorder, request)
	
	response := recorder.Result()
	defer response.Body.Close()
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	responseBody := map[string]any{}
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestGetAllCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepository()
	tx, _ := db.Begin()
	category1 := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	category2 := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Smartphone",
	})
	tx.Commit()

	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-KEY", "SECRET")

	router.ServeHTTP(recorder, request)
	
	response := recorder.Result()
	defer response.Body.Close()
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	responseBody := map[string]any{}
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	categories := responseBody["data"].([]any)
	categoryResponse1 := categories[0].(map[string]any)
	categoryResponse2 := categories[1].(map[string]any)

	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Name, categoryResponse1["name"])

	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoryResponse2["name"])
}

// func TestGetAllCategoryFailed(t *testing.T) {}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepository()
	tx, _ := db.Begin()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	requestBody := strings.NewReader(`{"name": "Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, 
	"http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "SECRET")

	router.ServeHTTP(recorder, request)
	
	response := recorder.Result()
	defer response.Body.Close()
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	responseBody := map[string]any{}
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]any)["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]any)["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepository()
	tx, _ := db.Begin()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPut, 
	"http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "SECRET")

	router.ServeHTTP(recorder, request)
	
	response := recorder.Result()
	defer response.Body.Close()
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	responseBody := map[string]any{}
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepository()
	tx, _ := db.Begin()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, 
	"http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-KEY", "SECRET")

	router.ServeHTTP(recorder, request)
	
	response := recorder.Result()
	defer response.Body.Close()
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	responseBody := map[string]any{}
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("X-API-KEY", "SECRET")

	router.ServeHTTP(recorder, request)
	
	response := recorder.Result()
	defer response.Body.Close()
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	responseBody := map[string]any{}
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-KEY", "WRONG")

	router.ServeHTTP(recorder, request)
	
	response := recorder.Result()
	defer response.Body.Close()
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	responseBody := map[string]any{}
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 401, response.StatusCode)
	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}