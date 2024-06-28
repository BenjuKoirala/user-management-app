package routes

import (
	"user-management-backend/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/api/users", controllers.GetUsers)
	e.POST("/api/users", controllers.CreateUser)
	e.PUT("/api/users/:id", controllers.UpdateUser)
	e.DELETE("/api/users/:id", controllers.DeleteUser)
}
