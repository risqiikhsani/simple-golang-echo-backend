package main

import (
	"myapp/handlers"
	"github.com/labstack/echo/v4"
)



func main() {
	// Create a new instance of Echo
	e := echo.New()

	// Define our routes
	e.GET("/items", handlers.GetItems)
	e.POST("/items", handlers.CreateItem)
	e.GET("/items/:id", handlers.GetItem)
	e.PUT("/items/:id", handlers.UpdateItem)
	e.DELETE("/items/:id", handlers.DeleteItem)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}

