package models

import "gorm.io/gorm"

type JobListing struct {
	gorm.Model
	ID          string  `json:"id"`
	Position    string  `json:"position"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Pay         float64 `json:"pay"`
	Company     string  `json:"company"`
	Salaried    bool    `json:"salaried"`
	Remote      bool    `json:"remote"`
}
