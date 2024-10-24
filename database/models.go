package database

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Id    int `gorm:"primarykey"`
	Name  string
	Group string
	Email string
}

type Subject struct {
	gorm.Model
	Id   int `gorm:"primaryKey"`
	Name string
}

type Grade struct {
	gorm.Model
	GradeID   int `gorm:"primaryKey"`
	StudentID int `gorm:"foreignKey"`
	SubjectID int `gorm:"foreignKey"`
	Grade     float64
}
