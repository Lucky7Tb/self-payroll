package main

import (
	"os"
	"self-payroll/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dotenvError := godotenv.Load()
	if dotenvError != nil {
		panic("Error loading .env")
	}

	router := echo.New()
	db := connectToDb()
	router.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	routes.AppRoute(router, db)
	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}

func connectToDb() *gorm.DB {
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT")
	connection, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("Can't connect to database")
	}

	return connection
}
