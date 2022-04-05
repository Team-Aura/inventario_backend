package models

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConexionABBDD() {
	dsn := "Cadena de conexion"
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Fallo en la conexion, abortando")
	}
	database.AutoMigrate(&Usuarios{})
	database.AutoMigrate(&UsuarioToken{})
	database.AutoMigrate(&Articulos{})
	database.AutoMigrate(&Almacenes{})
	database.AutoMigrate(&AlmacenAlmacenArticulosCantidades{})
	DB = database
}
