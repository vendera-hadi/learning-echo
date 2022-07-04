package main

import (
	"log"

	"learnecho/database"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Define API wrapper
	api := echo.New()
	db := database.GetInstance()
	m := database.GetMigrations(db)
	err = m.Migrate()
	if err == nil {
		print("Migrations did run successfully")
	} else {
		print("migrations failed.", err)
	}

	server := echo.New()
	server.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		if req.URL.Path[:4] == "/api" {
			api.ServeHTTP(res, req)
		}

		return
	})
	server.Logger.Fatal(server.Start(":1200"))
}
