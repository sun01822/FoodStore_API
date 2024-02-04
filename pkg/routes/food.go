package routes

import (
	"Food_API/pkg/controllers"
	"github.com/labstack/echo/v4"
)

type foodRoutes struct {
	echo *echo.Echo
	foodController controllers.IFoodController
}


func FoodRoutes(e *echo.Echo, foodController controllers.IFoodController) *foodRoutes {
return &foodRoutes{
		echo: e,
		foodController: foodController,
	}
}

func (fc *foodRoutes) InitFoodRoutes(){
	e := fc.echo
	fc.initFoodRoutes(e)
}

func (fc *foodRoutes) initFoodRoutes(e *echo.Echo){

	// group the routes
	food := e.Group("/foodstore")
	
	// Initializing http methods of Food - routing endpoints and calling the controller methods
	food.POST("/food", fc.foodController.CreateFood)
	food.GET("/food", fc.foodController.GetFoods)
	food.PUT("/food/:id", fc.foodController.UpdateFood)
	food.DELETE("/food/:id", fc.foodController.DeleteFood)
	food.GET("/food/:id", fc.foodController.GetFoodByID)

	// Search by category
	food.GET("/food/category/:category", fc.foodController.SearchFoodByCategory)

	// Sort by price
	food.GET("/food/sort", fc.foodController.SortFoodByPrice)

}