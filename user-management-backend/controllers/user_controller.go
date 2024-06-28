package controllers

import (
	"github.com/labstack/gommon/log"
	"net/http"
	"user-management-backend/models"
	"user-management-backend/services"

	"github.com/labstack/echo/v4"
)

var (
	getUsers   = services.GetUsers
	createUser = services.CreateUser
	updateUser = services.UpdateUser
	deleteUser = services.DeleteUser
)

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c echo.Context) error {
	log.Infof("Getting user information")
	users, err := getUsers()
	if err != nil {
		log.Errorf("Error while obtaining user info", err.Error())
		return c.JSON(http.StatusInternalServerError, "Can not get user details")
	}
	return c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} models.User
// @Router /users [post]
func CreateUser(c echo.Context) error {
	log.Info("Creating user")
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := createUser(user); err != nil {
		log.Errorf("Error while creating user info", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update details of an existing user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User Data"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		log.Errorf("Error while updating user info", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := updateUser(id, user); err != nil {
		log.Errorf("Error while updating user info", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if err := deleteUser(id); err != nil {
		log.Errorf("Error while deleting user info", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
