package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"user-management-backend/database"
	"user-management-backend/routes"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title User Management API
// @version 1.0
// @description This is a sample server for managing users.
// @host localhost:8080
// @BasePath /api
func main() {
	e := echo.New()

	config, err := loadConfig("config/db_config.json")
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	frontendUrl := config.FrontEndUrl

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{frontendUrl}, //angular service address
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Connect to the database
	database.ConnectDB()

	// Initialize routes
	routes.InitRoutes(e)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type Config struct {
	FrontEndUrl string `json:"frontend_url"`
}

// LoadConfig reads the configuration from a JSON file
func loadConfig(filePath string) (*Config, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	config := &Config{}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
