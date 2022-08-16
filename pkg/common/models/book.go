package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model // adds ID, CreatedAt, etc
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}