package database

type Student struct {
	Student_id int    `gorm:"primarykey"`
	Name       string `json:"name"`
	Group      string `json:"group"`
	Email      string `json:"email"`
}

type Subject struct {
	SubjectID uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
}

type Grade struct {
	GradeID   int     `gorm:"primaryKey"`
	StudentID int     `json:"student_id"`
	SubjectID int     `json:"subject_id"`
	Grade     float64 `json:"grade"`
}
