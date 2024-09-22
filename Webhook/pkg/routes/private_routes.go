package routes

import (
	"github.com/gofiber/fiber/v2"
	"simple.webhook/app/controllers"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/book", controllers.CreateBook) // create a new book

	// Routes for PUT method:
	route.Put("/book", controllers.UpdateBook) // update one book by ID

	// Routes for DELETE method:
	route.Delete("/book", controllers.DeleteBook) // delete one book by ID
}
