package models

type Articulos struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Nombre string `json:"nombre"`
	Codigo string `json:"codigo"`
}

type Almacenes struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Concepto  string `json:"concepto"`
	UsuarioID int
	Usuario   Usuarios `gorm:"References:Id"`
	Activo    bool     `json:"activo"`
}

type AlmacenArticulosCantidades struct {
	Id         uint `json:"id" gorm:"primary_key"`
	AlmacenID  int
	Almacen    Almacenes `gorm:"References:Id"`
	ArticuloID int
	Articulo   Articulos `gorm:"References:Id"`
	Cantidad   float32   `json:"cantidad"`
}
