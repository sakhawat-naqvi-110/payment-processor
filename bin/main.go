package main

import (
	"github.com/go-playground/validator/v10"
	"go/payment-processor/pkg/config"
	handler "go/payment-processor/pkg/handler"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Initialize Logger
	config.InitializeLogger()
	log := config.GetLogger()
	err := godotenv.Load()
	if err != nil {
		log.Info("Error loading .env file")
	}
	config.ConnectDB()

	// Initialize Echo
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	priv := e.Group("/payment-process")

	db := config.GetDb()

	handler.RegisterRoutes(priv, db, log, validator.New())

	log.Info("Server starting on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
