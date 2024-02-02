package routes

import (
	"Food_API/pkg/controllers"
	"fmt"

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
	food.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println(c.Request().Body)
			return next(c)
		}
	})


	// Initializing http methods of Food - routing endpoints and calling the controller methods
	food.POST("/food", fc.foodController.CreateFood)
	food.GET("/food", fc.foodController.GetFoods)
	food.PUT("/food/:id", fc.foodController.UpdateFood)
	food.DELETE("/food/:id", fc.foodController.DeleteFood)

	// Search by category
	food.GET("/food/category/:category", fc.foodController.SearchFoodByCategory)

}