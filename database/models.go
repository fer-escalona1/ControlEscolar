package database

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name  string
	Group string
	Email string
}

type Subject struct {
	gorm.Model
	Name string
}

type Grade struct {
	gorm.Model
	StudentID int `gorm:"foreignKey"`
	SubjectID int `gorm:"foreignKey"`
	Grade     float64
}
