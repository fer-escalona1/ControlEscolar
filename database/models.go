package database

type Student struct{
	Student_id int `gorm:"primarykey"`
	Name string
	Group string
	Email string
}

type Subject struct{
	Subject_id int `gorm:"primarykey"`
	Name string
}