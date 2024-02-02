package services

import (
	"Food_API/pkg/domain"
	"Food_API/pkg/models"
	"errors"
	"gorm.io/gorm"
)

// Parent struct to implement interface binding
type foodService struct {
	repo domain.IFoodRepo
}

func FoodServiceInstance(userRepo domain.IFoodRepo) domain.IFoodService {
	return &foodService{
		repo: userRepo,
	}
}

// GetFoods implements domain.IFoodService.
func (service *foodService) GetFoods(model *gorm.Model) ([]models.Food, error) {
	var allFoods []models.Food
	food, _ := service.repo.GetFoods(model)
	if len(food) == 0{
		return nil, errors.New("No Food Item Found")
	}
	allFoods = append(allFoods, food...)
	return allFoods, nil
}



// CreateUser implements domain.IFoodService.
func (service *foodService) CreateFood(food *models.Food) error {
	if err := service.repo.CreateFood(food); err != nil {
		return err
	}
	return nil
}

// DeleteUser implements domain.IFoodService.
func (service *foodService) DeleteFood(model *gorm.Model) error {
	if err := service.repo.DeleteFood(model); err != nil {
		return errors.New("Food Item is not deleted")
	}
	return nil
}


// UpdateUser implements domain.IFoodService.
func (service *foodService) UpdateFood(food *models.Food) error {
	if err := service.repo.UpdateFood(food); err != nil {
		return errors.New("Food Item is not updated")
	}
	return nil
}

// SearchByCategory implements domain.IFoodService.
func (service *foodService) SearchByCategory(category string) ([]models.Food, error) {
	var allFoods []models.Food
	food, _ := service.repo.SearchByCategory(category)
	if len(food) == 0{
		return nil, errors.New("No Food Item Found")
	}
	allFoods = append(allFoods, food...)
	return allFoods, nil
}
