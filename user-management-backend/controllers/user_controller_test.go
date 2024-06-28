package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"user-management-backend/models"
)

func TestUserController(t *testing.T) {
	// Initializing Echo
	e := echo.New()

	t.Run("GetUsers", func(t *testing.T) {
		// Mock GetUsers service function
		mockUsers := []models.User{
			{ID: 1, Name: "John", Email: "john@example.com"},
			{ID: 2, Name: "Alice", Email: "alice@example.com"},
		}
		getUsers = func() ([]models.User, error) {
			return mockUsers, nil
		}

		// Creating a new HTTP request
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Calling GetUsers controller function
		if assert.NoError(t, GetUsers(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			// Assert response body
			expectedBody := `[{"id":1,"name":"John","email":"john@example.com"},{"id":2,"name":"Alice","email":"alice@example.com"}]`
			assert.JSONEq(t, expectedBody, rec.Body.String())
		}
	})

	t.Run("CreateUser", func(t *testing.T) {
		// Mocking CreateUser service function
		createUser = func(user *models.User) error {
			user.ID = 1 // Simulate creation of user
			return nil
		}

		// Preparing request body
		requestBody := `{"name":"Alice","email":"alice@example.com"}`

		// Creating a new HTTP request
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Calling CreateUser controller function and checking
		if assert.NoError(t, CreateUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			expectedBody := `{"id":1,"name":"Alice","email":"alice@example.com"}`
			assert.JSONEq(t, expectedBody, rec.Body.String())
		}
	})

	t.Run("UpdateUser", func(t *testing.T) {
		// Mocking UpdateUser service function
		updateUser = func(id string, user *models.User) error {
			user.ID = 1
			return nil
		}

		// Preparing request body
		requestBody := `{"name":"Alice","email":"alice@example.com"}`

		// Creating a new HTTP request
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		// Calling UpdateUser controller function
		if assert.NoError(t, UpdateUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			expectedBody := `{"id":1,"name":"Alice","email":"alice@example.com"}`
			assert.JSONEq(t, expectedBody, rec.Body.String())
		}
	})

	t.Run("DeleteUser", func(t *testing.T) {
		// Mocking DeleteUser service function
		deleteUser = func(id string) error {
			return nil
		}

		// Creating a new HTTP request
		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		// Calling DeleteUser controller function
		if assert.NoError(t, DeleteUser(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	})
}
