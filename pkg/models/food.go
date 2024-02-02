package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Calories string  `json:"calories"`
	Price    float64 `json:"price"`
	Details  string  `json:"details"`
	AvailableTime string `json:"available_time"`
	IsAvailable string `json:"is_available" gorm:"default:Available"`
}