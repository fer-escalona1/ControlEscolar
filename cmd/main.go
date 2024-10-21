package main

import (
	"fmt"
	"net/http"

	"github.com/Amari05fc/ControlEscolar/database"
	"github.com/gin-gonic/gin"
)

type Student struct{
	Student_id int `json: "student_id"`
	Name string `json: "name"`
	Group string `json: "group"`
	Email string `json: "email"`
}

type Subject struct{
	Subject_id int `json: "subject_id"`
	Name string `json: "name"`
}

func main() {
	db, err := database.NewDatabaseDriver()
	if err != nil{
		fmt.Println("Error al conectar a la base de datos: DB_estudiantes", err)
		return
	}

	fmt.Println("BD")
	db.AutoMigrate(&database.Student{})
	db.AutoMigrate(&database.Subject{})
	db.AutoMigrate(&database.Grade{})
	fmt.Println("AutoMigrate")

	r := gin.Default()
	students := []Student{}
	subjects := []Subject{}
	indexStudent := 1
	//indexSubject := 1

	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message" : "pong",
		})		
	})

	//API GET Estudiante
	r.GET("/api/students", func(c *gin.Context) {
		db.Find(&students)
		c.JSON(200, students)
	})

		//Get un solo estudiante
		r.GET("/api/students/:student_id", func(c *gin.Context) {
			stu_id := c.Param("student_id")
			fmt.Println("GET estudiante", stu_id)
			db.Find(&students, stu_id)
			c.JSON(200, students)
			db.Find(stu_id)
		})

	//API GET Materia
	r.GET("/api/subjects/:subject_id", func(c *gin.Context) {
		sub_id := c.Param("subject_id")
		fmt.Println("GET materia", sub_id)
		db.Find(&subjects, sub_id)
		c.JSON(200, subjects)
		db.Get(sub_id)
	})

	//Creamos estudiante POST
	r.POST("/api/students", func(c *gin.Context) {
		var student Student

		db.Last(&student)
		indexStudent = student.Student_id
		indexStudent++

		if c.BindJSON(&student) == nil {
			student.Student_id = indexStudent

			c.JSON(200, student)
			db.Create(&student)
			fmt.Println("Se creo un estudiante", student.Student_id)
		}
	})

	r.Run(":8001")

}