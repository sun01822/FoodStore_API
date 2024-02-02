package controllers

import (
	"Food_API/pkg/domain"
	"Food_API/pkg/models"
	"Food_API/pkg/types"
	"net/http"
	"strconv"
	"time"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type IFoodController interface {
	GetFoods(c echo.Context) error
	GetFoodByID(c echo.Context) error
	CreateFood(c echo.Context) error
	UpdateFood(c echo.Context) error
	DeleteFood(c echo.Context) error
	SearchFoodByCategory(c echo.Context) error
}

type foodController struct {
	foodSvc domain.IFoodService
}

func FoodControllerInstance(foodSvc domain.IFoodService) IFoodController {
	return &foodController{
		foodSvc: foodSvc,
	}
}


// CreateFood implements IFoodController.
func (controller *foodController) CreateFood(e echo.Context) error {
	reqUser := &types.FoodRequest{}
	fmt.Println(reqUser)
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid data request")
	}
	if err := reqUser.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	food := &models.Food{
		Name:          reqUser.Name,
		Category:      reqUser.Category,
		Calories:      reqUser.Calories,
		Price:         reqUser.Price,
		Details:       reqUser.Details,
		AvailableTime: reqUser.AvailableTime,
		IsAvailable:   reqUser.IsAvailable,
	}
	fmt.Println(food)
	if err := controller.foodSvc.CreateFood(food); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "Food created successfully")
}

// GetFoods implements IFoodController.
func (controller *foodController) GetFoods(e echo.Context) error {
	// tempFoodID := e.QueryParam("id")
	// FoodID, err := strconv.ParseInt(tempFoodID, 0, 0)
	// if err != nil && tempFoodID != "" {
	// 	return e.JSON(http.StatusBadRequest, "Enter a valid Food ID")
	// }
	Food, err := controller.foodSvc.GetFoods(&gorm.Model{})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, Food)
}

// GetFoodByID implements IFoodController.
func (controller *foodController) GetFoodByID(e echo.Context) error {
	tempFoodID := e.Param("id")
	FoodID, err := strconv.ParseInt(tempFoodID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid Food ID")
	}	
	Food, err := controller.foodSvc.GetFoodByID(&gorm.Model{ID: uint(FoodID)})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	if Food.ID == 0 {
		return e.JSON(http.StatusNotFound, "Food Item is not found")
	}
	return e.JSON(http.StatusOK, Food)
}


// DeleteFood implements IFoodController.
func (controller *foodController) DeleteFood(e echo.Context) error {
	tempFoodID := e.Param("id")
	FoodID, err := strconv.ParseInt(tempFoodID, 0, 0)
	if err != nil && tempFoodID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid Food ID")
	}
	_, err = controller.foodSvc.GetFoods(&gorm.Model{ID: uint(FoodID)})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := controller.foodSvc.DeleteFood(&gorm.Model{ID: uint(FoodID)}); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Food deleted successfully") 
}


// UpdateFood implements IFoodController.
func (controller *foodController) UpdateFood(e echo.Context) error {
	reqFood := &types.FoodRequest{}
	if err := e.Bind(reqFood); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	tempFoodID := e.Param("id")
	FoodID, err := strconv.ParseInt(tempFoodID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid Food ID")
	}
	existingFood, err := controller.foodSvc.GetFoods(&gorm.Model{ID: uint(FoodID)})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	updateFood := &models.Food{
		Model:     gorm.Model{ID: uint(FoodID), UpdatedAt: time.Now(), CreatedAt: existingFood[0].CreatedAt, DeletedAt: existingFood[0].DeletedAt},
		Name:          reqFood.Name,
		Category:      reqFood.Category,
		Calories:      reqFood.Calories,
		Price:         reqFood.Price,
		Details:       reqFood.Details,
		AvailableTime: reqFood.AvailableTime,
		IsAvailable:   reqFood.IsAvailable,
	}
	if updateFood.Name == "" {
		updateFood.Name = existingFood[0].Name
	}
	if updateFood.Category == "" {
		updateFood.Category = existingFood[0].Category
	}
	if updateFood.Calories == "" {
		updateFood.Calories = existingFood[0].Calories
	}
	if updateFood.Price == 0 {
		updateFood.Price = existingFood[0].Price
	}
	if updateFood.Details == "" {
		updateFood.Details = existingFood[0].Details
	}
	if updateFood.AvailableTime == "" {
		updateFood.AvailableTime = existingFood[0].AvailableTime
	}
	if updateFood.IsAvailable == "" {
		updateFood.IsAvailable = existingFood[0].IsAvailable
	}
	if err := controller.foodSvc.UpdateFood(updateFood); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Food updated successfully")
}

// 

// SearchFoodByCategory implements IFoodController.
func (controller *foodController) SearchFoodByCategory(e echo.Context) error {
	tempCategory := e.Param("category")
	if tempCategory == "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid category")
	}
	Food, err := controller.foodSvc.SearchByCategory(tempCategory)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, Food)
}