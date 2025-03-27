package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rodrigoferra/ports-blog/internal/core/adapters"
	"github.com/rodrigoferra/ports-blog/internal/infrastructure"
	"github.com/rodrigoferra/ports-blog/internal/infrastructure/handlers"
	"github.com/rodrigoferra/ports-blog/internal/infrastructure/repositories"
	"log"
)

func main() {
	// Initialize database
	db := infrastructure.InitPostgresDatabase(
		"localhost", 5432, "postgres", "admin", "postgres",
	)

	// Initialize dependencies
	postRepo := repositories.NewGormPostRepository(db)
	postService := adapters.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Custom validator
	e.Validator = infrastructure.NewCustomValidator()

	// Routes
	e.POST("/posts", postHandler.Create)
	e.GET("/posts", postHandler.List)
	e.GET("/posts/:id", postHandler.GetByID)
	e.PUT("/posts/:id", postHandler.Update)
	e.DELETE("/posts/:id", postHandler.Delete)

	// Start server
	log.Fatal(e.Start(":8080"))
}
