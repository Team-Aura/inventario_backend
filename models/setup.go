package models

import (
	"TeamAura/go/ORM/utils"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConexionABBDD() {
	dsn := utils.Uri
	var err error = nil
	var database *gorm.DB = nil
	if utils.Usuario == "Berna" {
		database, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	} else {
		database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	}

	if err != nil {
		panic("Fallo en la conexion, abortando")
	}
	database.AutoMigrate(&Usuarios{})
	database.AutoMigrate(&UsuarioToken{})
	database.AutoMigrate(&Articulos{})
	database.AutoMigrate(&Almacenes{})
	database.AutoMigrate(&AlmacenArticulosCantidades{})
	DB = database
}
