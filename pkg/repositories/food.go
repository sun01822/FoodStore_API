package repositories

import (
	"Food_API/pkg/domain"
	"Food_API/pkg/models"
	"gorm.io/gorm"
)

// Parent struct to implement interface binding
type foodRepo struct {
	d *gorm.DB
}

// Interface binding
func FoodDBInstance(d *gorm.DB) domain.IFoodRepo {
	return &foodRepo{
		d: d,
	}
}


// GetFoods function to get all foods from database
func (repo *foodRepo) GetFoods(model *gorm.Model) ([]models.Food, error) {
	var foods []models.Food
	// var err error
	// foodID := model.ID
	// if foodID != 0 {
	// 	err = repo.d.Where("id = ?", foodID).Find(&foods).Error
	// } else {
	// 	err = repo.d.Find(&foods).Error
	// }
	// if err != nil {
	// 	return []models.Food{}, err
	// }
	// return foods, nil
	err := repo.d.Find(&foods).Error
	if err != nil {
		return []models.Food{}, err
	}
	return foods, nil
}

// GetFoodByID function to get a food by ID from database
func (repo *foodRepo) GetFoodByID(model *gorm.Model) (models.Food, error) {
	var food models.Food
	err := repo.d.Where("id = ?", model.ID).Find(&food).Error
	if err != nil {
		return models.Food{}, err
	}
	return food, nil
}

// CreateFood function to create a new food in database
func (repo *foodRepo) CreateFood(food *models.Food) error {
	err := repo.d.Create(&food).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateFood function to update a food in database
func (repo *foodRepo) UpdateFood(food *models.Food) error {
	err := repo.d.Save(&food).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteFood function to delete a food from database
func (repo *foodRepo) DeleteFood(model *gorm.Model) error {
	var food models.Food
	err := repo.d.Where("id = ?", model.ID).Delete(&food).Error
	if err != nil {
		return err
	}
	return nil
}

// SearchByCategory function to search food by category
func (repo *foodRepo) SearchByCategory(category string) ([]models.Food, error) {
	var foods []models.Food
	err := repo.d.Where("category=?", category).Find(&foods).Error
	if err != nil {
		return []models.Food{}, err
	}
	return foods, nil
}
