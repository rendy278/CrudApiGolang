package main

import (
	"gocrud/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userController := controllers.NewUserController()

	// Routes
	e.POST("/users", userController.CreateUser)
	e.GET("/users", userController.GetUsers)
	e.GET("/users/:id", userController.GetUser)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
