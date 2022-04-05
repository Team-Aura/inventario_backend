package models

type Usuarios struct {
	Id            uint   `json:"id" gorm:"primary_key"`
	NombreUsuario string `json:"nombre"`
	Codigo        string `json:"codigo"`
	Contrasenia   string `json:"contrasenia"`
	Activo        bool   `json:"activo"`
}

type UsuarioToken struct {
	UsuarioId uint     `json:"UsuarioId" gorm:"primary_key"`
	Usuario   Usuarios `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Token     string   `json:"token"`
}
