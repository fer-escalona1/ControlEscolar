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
	fmt.Println("AutoMigrate")

	r := gin.Default()
	//students := []Student{}
	//subjects := []Subject{}

	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message" : "pong",
		})		
	})


	r.Run(":8001")

}