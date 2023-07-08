package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title string
	Genre string
}
