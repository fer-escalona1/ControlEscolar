package database

import (
	"fmt"
	"github.com/Amari05fc/ControlEscolar/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseDriver() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.MYSQL_DATABASE_URL), &gorm.Config{})
	if err != nil{
		fmt.Println("Error al conectar a la base de datos: BD_escolar", err)
	}
	return db, nil
}