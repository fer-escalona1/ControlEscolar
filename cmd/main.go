package main

import (
	"fmt"
	"net/http"
	// "strconv"

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

type Grade struct {
	GradeID   int     `json: "grade_id"`
	StudentID int     `json: "student_id"`
	SubjectID int     `json: "subject_id"`
	Grade     float64 `json:"grade"`
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
	//db.AutoMigrate(&database.Grade{})
	fmt.Println("AutoMigrate")

	r := gin.Default()
	students := []Student{}
	subjects := []Subject{}
	indexStudent := 1
	indexSubject := 1

	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message" : "pong",
		})		
	})

//-------------------------STUDENTS--------------------------------

	//API GET Estudiante
	// /api/strudents
	r.GET("/api/students", func(c *gin.Context) {
		db.Find(&students)
		c.JSON(200, students)
	})

		//Get un solo estudiante
		// /api/students/:student_id
		r.GET("/api/students/:student_id", func(c *gin.Context) {
			stu_id := c.Param("student_id")
			fmt.Println("GET estudiante", stu_id)
			db.Find(&students, stu_id)
			c.JSON(200, students)
			db.Find(stu_id)
		})

	//Creamos estudiante POST
	// /api/students
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

	//	Eliminamos un estudiante DELETE
	// /api/students/:student_id
	r.DELETE("/api/students/:student_id", func(c *gin.Context) {
		var student Student
		id := c.Param("student_id")
		if err != nil {
			c.JSON(400, gin.H{
				"erro": "Invalid ID",
			})
			return
		}
		db.Delete(student, id)
		c.JSON(201, gin.H{})
	})

	//	Actualizamos un estudiante PUT    PENDIENTE
	// /api/students/:student_id
	/*
	r.PUT("/api/students/:student_id", func(c *gin.Context) {
		var student Student

		id := c.Param("student_id")
		fmt.Println(id)
		idParseEstudiante, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid id",
			})
			return
		}

		err = c.BindJSON(&student)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid payload",
			})
			return
		}

		student.Student_id = idParseEstudiante
		//db.First(&student)
		fmt.Println("Busqueda estudiante", student.Email)
		
		fmt.Println("ID estudiante: ",student.Student_id)
		db.Save(&student)
		
		c.JSON(201, gin.H{})
	})
	*/

//-------------------------SUBJECTS--------------------------------

	//API GET Materia
	// /api/subjects/:subject_id
	r.GET("/api/subjects/:subject_id", func(c *gin.Context) {
		sub_id := c.Param("subject_id")
		fmt.Println("GET materia", sub_id)
		db.Find(&subjects, sub_id)
		c.JSON(200, subjects)
		db.Get(sub_id)
	})

	//Crear Materias POST
	// /api/subjects
	r.POST("/api/subjects", func(c *gin.Context) {
		var subjetc Subject

		db.Last(&subjetc)
		indexSubject = subjetc.Subject_id
		indexSubject++

		if c.BindJSON(&subjetc) == nil{
			subjetc.Subject_id = indexSubject

			c.JSON(200, subjetc)
			db.Create(&subjetc)
			fmt.Println("Se creo una Materia", subjetc.Subject_id)
		}
	})

	//Eliminar Materias DELETE
	// /api/subjects/:subject_id
	r.DELETE("/api/subjects/:subject_id", func(c *gin.Context) {
		var subject Subject
		id := c.Param("subject_id")
		if err != nil {
			c.JSON(400, gin.H{
				"erro": "Invalid ID",
			})
			return
		}
		db.Delete(subject, id)
		c.JSON(201, gin.H{})
	})

	//Actualizamos Materias PUT
	// /api/subjects
	/*
	r.PUT("/api/subjects", func(c *gin.Context) {
		var subject Subject
		id := c.Param("subject_id")
		idParsed, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid id",
			})
			return
		}

		err = c.BindJSON(&subject)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid payload",
			})
			return
		}

		subject.Subject_id = idParsed
		db.Save(&subject)
	})
	*/

//-------------------------GRADE--------------------------------

	//Obtenemos calificaciones de materia de un estudiante GET
	// /api/grades/:grade_id/student/:student_id
	r.GET("/api/grades/:grade_id/student/:student_id", func(c *gin.Context) {})

	//Obtener todas las calificaciones de un estudiante GET
	// /api/grades/student/:studen_id
	r.GET("/api/grades/student/:studen_id", func(c *gin.Context) {})

	//Crear calificación POST
	// /api/grades
	r.POST("/api/grades", func(c *gin.Context) {
		var grade Grade

		db.Last(&grade)
		indexSubject = grade.GradeID
		indexSubject++

		if c.BindJSON(&grade) == nil{
			grade.GradeID = indexSubject

			c.JSON(200, grade)
			db.Create(&grade)
			fmt.Println("Se creo una Materia", grade.GradeID)
		}
	})

	//Eliminar calificación DELETE
	// /api/grades/:grade_id
	r.DELETE("/api/grades/:grade_id", func(c *gin.Context) {
		var grade Grade
		id := c.Param("subject_id")
		if err != nil {
			c.JSON(400, gin.H{
				"erro": "Invalid ID",
			})
			return
		}
		db.Delete(grade, id)
		c.JSON(201, gin.H{})
	})

	//Actualizar calificación PUT
	// /api/grades/:grade_id
	r.PUT("/api/grades/:grade_id", func(c *gin.Context) { })


	r.Run(":8001")

}