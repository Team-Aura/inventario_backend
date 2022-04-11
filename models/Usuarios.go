package models

import "gorm.io/gorm"

type Usuarios struct {
	gorm.Model
	Id            uint   `json:"-" gorm:"primary_key"`
	NombreUsuario string `json:"nombre"`
	Codigo        string `json:"codigo"`
	Contrasenia   string `json:"-"`
	Activo        bool   `json:"activo"`
}

type UsuarioToken struct {
	gorm.Model
	UsuarioID int
	Usuario   Usuarios `gorm:"References:Id"`
	Token     string   `json:"token" gorm:"primary_key"`
}
