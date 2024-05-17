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
	// Mock services and setup Echo instance
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("GetUsers", func(t *testing.T) {
		// Mock GetUsers service function
		mockUsers := []models.User{
			{ID: 1, Name: "John", Email: "john@example.com"},
			{ID: 2, Name: "Alice", Email: "alice@example.com"},
		}
		getUsers = func() ([]models.User, error) {
			return mockUsers, nil
		}

		// Call GetUsers controller function
		if assert.NoError(t, GetUsers(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			// Assert response body
			expectedBody := `[{"id":1,"name":"John","email":"john@example.com"},{"id":2,"name":"Alice","email":"alice@example.com"}]`
			assert.JSONEq(t, expectedBody, rec.Body.String())
		}
	})

	t.Run("GetUser", func(t *testing.T) {
		// Mock GetUser service function
		mockUser := models.User{ID: 1, Name: "John", Email: "john@example.com"}
		getUser = func(id string) (models.User, error) {
			return mockUser, nil
		}

		// Call GetUser controller function
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		if assert.NoError(t, GetUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			// Assert response body
			expectedBody := `{"id":1,"name":"John","email":"john@example.com"}`
			assert.JSONEq(t, expectedBody, rec.Body.String())
		}
	})

	t.Run("CreateUser", func(t *testing.T) {
		// Mock CreateUser service function
		//mockUser := models.User{ID: 1, Name: "Alice", Email: "alice@example.com"}
		createUser = func(user *models.User) error {
			user.ID = 1 // Simulate creation of user
			return nil
		}

		// Prepare request body
		requestBody := `{"name":"Alice","email":"alice@example.com"}`

		// Set request body
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call CreateUser controller function
		if assert.NoError(t, CreateUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			// Assert response body
			expectedBody := `{"id":1,"name":"Alice","email":"alice@example.com"}`
			assert.JSONEq(t, expectedBody, rec.Body.String())
		}
	})

	t.Run("UpdateUser", func(t *testing.T) {
		// Mock UpdateUser service function
		updateUser = func(id string, user *models.User) error {
			// Simulate update of user
			return nil
		}

		// Prepare request body
		requestBody := `{"name":"Alice","email":"alice@example.com"}`

		// Set request body
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		// Call UpdateUser controller function
		if assert.NoError(t, UpdateUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			// Assert response body
			expectedBody := `{"id":1,"name":"Alice","email":"alice@example.com"}`
			assert.JSONEq(t, expectedBody, rec.Body.String())
		}
	})

	t.Run("DeleteUser", func(t *testing.T) {
		// Mock DeleteUser service function
		deleteUser = func(id string) error {
			// Simulate deletion of user
			return nil
		}

		// Call DeleteUser controller function
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		if assert.NoError(t, DeleteUser(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	})
}
