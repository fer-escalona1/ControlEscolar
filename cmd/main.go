package main

import (
	"fmt"

	"github.com/Amari05fc/ControlEscolar/database"
)

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
}