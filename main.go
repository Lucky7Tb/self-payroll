package main

import (
	"os"
	"self-payroll/config"
	"self-payroll/routes"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dotenvError := godotenv.Load()
	if dotenvError != nil {
		panic("Error loading .env")
	}

	router := echo.New()
	db := config.ConnectToDb()
	router.Validator = &config.Validator{Validator: validator.New()}
	router.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	routes.AppRoute(router, db)
	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
