package models

type Articulos struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Nombre string `json:"nombre"`
	Codigo string `json:"codigo"`
}

type Almacenes struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Concepto  string `json:"concepto"`
	Activo    bool   `json:"activo"`
	UsuarioId int
	Usuario   Usuarios `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type AlmacenAlmacenArticulosCantidades struct {
	Id         uint `json:"id" gorm:"primary_key"`
	AlmacenId  int
	Almacen    Almacenes `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ArticuloId int
	Articulo   Articulos `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cantidad   float32   `json:"cantidad"`
}
