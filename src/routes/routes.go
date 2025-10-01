package routes

import (
	"api-e-ticketing/src/controllers"
	middleware "api-e-ticketing/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r fiber.Router) {
	r.Post("/auth/login", controllers.Login)

	r.Use(middleware.JwtMiddleware())
	admin := r.Group("/admin", middleware.RequireRole("admin"))
	admin.Post("/terminal", controllers.CreateTerminal)
}