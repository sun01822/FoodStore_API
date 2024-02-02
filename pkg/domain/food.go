package domain

import (
	"Food_API/pkg/models"
	"gorm.io/gorm"
)

// for database Repository operation (call from service)
type IFoodRepo interface {
	GetFoods(*gorm.Model) ([]models.Food, error)
	CreateFood(food *models.Food) error
	UpdateFood(food *models.Food) error
	DeleteFood(*gorm.Model) error
	SearchByCategory(string) ([]models.Food, error)
}

// for service operation (call from controller)
type IFoodService interface{
	GetFoods(*gorm.Model) ([]models.Food, error)
	CreateFood(user *models.Food) error
	UpdateFood(user *models.Food) error
	DeleteFood(*gorm.Model) error
	SearchByCategory(string) ([]models.Food, error)
}