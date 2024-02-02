package containers

import (
	"Food_API/pkg/config"
	"Food_API/pkg/connection"
	"Food_API/pkg/controllers"
	"Food_API/pkg/repositories"
	"Food_API/pkg/routes"
	"Food_API/pkg/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

// Serve is a function that returns a new instance of echo.Echo
func Serve(e *echo.Echo)  {

	// Config initalization
	config.SetConfig()

	// Database initialization
	db:= connection.GetDB()

	// Repository initialization
	foodRepository := repositories.FoodDBInstance(db)

	// Service initialization
	foodService := services.FoodServiceInstance(foodRepository)

	// Controller initialization
	foodController := controllers.FoodControllerInstance(foodService)

	// Routes 
	food := routes.FoodRoutes(e, foodController)

	food.InitFoodRoutes()

	// Starting Server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}