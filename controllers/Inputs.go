package controllers

//Articulos

type ArticulosInput struct {
	Nombre string `json:"nombre" binding:"required"`
	Codigo string `json:"codigo" binding:"required"`
}
type AlmacenesInput struct {
	Concepto string `json:"concepto" binding:"required"`
	Activo   bool   `json:"activo"`
	Usuario  int    `json:"usuario" binding:"required"`
}
type AlmacenArticulosCantidadesInput struct {
	AlmacenId  int     `json:"almacen" binding:"required"`
	ArticuloId int     `json:"articulo" binding:"required"`
	Cantidad   float32 `json:"cantidad" binding:"required"`
}

//Usuarios

type UserInput struct {
	NombreUsuario string `json:"nombre" binding:"required"`
	Contrasenia   string `json:"contrasenia" binding:"required"`
}
