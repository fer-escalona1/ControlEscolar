package main

import (
	"net/http"

	"github.com/Amari05fc/ControlEscolar/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa el router de Gin
	r := gin.Default()

	// Carga las plantillas HTML
	r.LoadHTMLGlob("templates/*")

	// Ruta principal para cargar la página
	r.GET("/", func(c *gin.Context) {
		var students []database.Student
		var subjects []database.Subject
		var grades []database.Grade

		// Conexión a la base de datos
		db, err := database.NewDatabaseDriver()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error de conexión a la base de datos")
			return
		}

		// Obtiene los registros de estudiantes, materias y calificaciones
		db.Find(&students)
		db.Find(&subjects)
		db.Find(&grades)

		// Renderiza la plantilla HTML con los datos obtenidos
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "Control Escolar",
			"students": students,
			"subjects": subjects,
			"grades":   grades,
		})
	})

	// API routes
	r.POST("/api/students", createStudent)
	r.DELETE("/api/students/:id", deleteStudent)
	r.POST("/api/subjects", createSubject)
	r.DELETE("/api/subjects/:id", deleteSubject)
	r.POST("/api/grades", createGrade)
	r.DELETE("/api/grades/:id", deleteGrade)

	// Inicia el servidor en el puerto 8001
	r.Run(":8001")
}

// createStudent maneja la creación de estudiantes
func createStudent(c *gin.Context) {
	var student database.Student
	if err := c.ShouldBindJSON(&student); err == nil {
		db, _ := database.NewDatabaseDriver()
		db.Create(&student) // Crea un nuevo registro en la base de datos
		c.JSON(http.StatusCreated, student)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// deleteStudent maneja la eliminación de estudiantes
func deleteStudent(c *gin.Context) {
	id := c.Param("id")
	db, _ := database.NewDatabaseDriver()
	db.Delete(&database.Student{}, id) // Elimina el registro de la base de datos
	c.Status(http.StatusNoContent)
}

// createSubject maneja la creación de materias
func createSubject(c *gin.Context) {
	var subject database.Subject
	if err := c.ShouldBindJSON(&subject); err == nil {
		db, _ := database.NewDatabaseDriver()
		db.Create(&subject) // Crea un nuevo registro en la base de datos
		c.JSON(http.StatusCreated, subject)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// deleteSubject maneja la eliminación de materias
func deleteSubject(c *gin.Context) {
	id := c.Param("id")
	db, _ := database.NewDatabaseDriver()
	db.Delete(&database.Subject{}, id) // Elimina el registro de la base de datos
	c.Status(http.StatusNoContent)
}

// createGrade maneja la creación de calificaciones
func createGrade(c *gin.Context) {
	var grade database.Grade
	if err := c.ShouldBindJSON(&grade); err == nil {
		db, _ := database.NewDatabaseDriver()
		db.Create(&grade) // Crea un nuevo registro en la base de datos
		c.JSON(http.StatusCreated, grade)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// deleteGrade maneja la eliminación de calificaciones
func deleteGrade(c *gin.Context) {
	id := c.Param("id")
	db, _ := database.NewDatabaseDriver()
	db.Delete(&database.Grade{}, id) // Elimina el registro de la base de datos
	c.Status(http.StatusNoContent)
}
