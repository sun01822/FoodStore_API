package types

import (
	validate "github.com/go-ozzo/ozzo-validation"
)

// FoodRequest is a struct for food request
type FoodRequest struct {
	Name string `json:"name"`
	Category string `json:"category"`
	Calories string `json:"calories,omitempty"`
	Price float64 `json:"price"`
	Details string `json:"details,omitempty"`
	AvailableTime string `json:"available_time"`
	IsAvailable string `json:"is_available"`
}

func (food FoodRequest) Validate() error {
	return validate.ValidateStruct(
		&food, 
		validate.Field(&food.Name, validate.Required.Error("False")),
		validate.Field(&food.Category, validate.Required.Error("False")),
		validate.Field(&food.Price, validate.Required.Error("False")),
		validate.Field(&food.AvailableTime, validate.Required.Error("False")),
	)
}
