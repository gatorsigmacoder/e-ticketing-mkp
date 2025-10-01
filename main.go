package main

import (
	"api-e-ticketing/src/database"
	"api-e-ticketing/src/routes"
	"api-e-ticketing/src/utils"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't read environment")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3001" // Default port
	}
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, PATCH",
		AllowCredentials: true,
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"status":  "UP",
			"version": "1.0.0",
			"uptime":  time.Since(time.Now()).String(),
		}
		return utils.Success(c, fiber.StatusOK, "Ok", data, nil)
	})

	api := app.Group("/api/v1")
	database.DatabaseInit()
	routes.RouteInit(api)
	err = app.Listen(port)
	if err != nil {
		log.Fatal("Error on starting app : " + err.Error())
	}
}