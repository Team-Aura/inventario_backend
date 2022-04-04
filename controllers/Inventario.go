package controllers

import (
	"TeamAura/go/ORM/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SELECTS DE LAS DIFERENTES TABLAS

func ConseguirArticulos(c *gin.Context) {
	var articulos []models.Articulos
	models.DB.Find(&articulos)
	c.JSON(http.StatusOK, gin.H{"articulos": articulos})
}

func ConseguirAlmacenes(c *gin.Context) {
	var almacenes []models.Almacenes
	models.DB.Find(&almacenes)
	c.JSON(http.StatusOK, gin.H{"almacenes": almacenes})
}

func ConseguirAlmacenesCantidades(c *gin.Context) {
	var cantidades []models.AlmacenAlmacenArticulosCantidades
	models.DB.Find(&cantidades)
	c.JSON(http.StatusOK, gin.H{"AlmacenAlmacenArticulosCantidades": cantidades})
}

//INSERTS DE LAS DIFERENTES TABLAS

type CrearArticulosInput struct {
	Nombre string `json:"nombre" binding:"required"`
	Codigo string `json:"codigo" binding:"required"`
}
type CrearAlmacenesInput struct {
	Concepto string `json:"concepto" binding:"required"`
	Activo   bool   `json:"activo" binding:"required"`
	Usuario  int    `json:"usuario" binding:"required"`
}
type CrearAlmacenAlmacenArticulosCantidadesInput struct {
	AlmacenId  int     `json:"almacen" binding:"required"`
	ArticuloId int     `json:"articulo" binding:"required"`
	Cantidad   float32 `json:"cantidad" binding:"required"`
}

func CrearArticulos(c *gin.Context) {
	// Validamos los datos de entrada
	var input CrearArticulosInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Creamos el articulo
	articulo := models.Articulos{Nombre: input.Nombre, Codigo: input.Codigo}
	models.DB.Create(&articulo)

	c.JSON(http.StatusOK, gin.H{"resultado": articulo})
}

func CrearAlmacenes(c *gin.Context) {
	// Validamos los datos de entrada
	var input CrearAlmacenesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Creamos el articulo
	almacen := models.Almacenes{Concepto: input.Concepto, Activo: input.Activo, UsuarioId: input.Usuario}
	models.DB.Create(&almacen)

	c.JSON(http.StatusOK, gin.H{"resultado": almacen})
}

func CrearAlmacenAlmacenArticulosCantidades(c *gin.Context) {
	// Validamos los datos de entrada
	var input CrearAlmacenAlmacenArticulosCantidadesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Creamos el articulo
	cantidad := models.AlmacenAlmacenArticulosCantidades{AlmacenId: input.AlmacenId, ArticuloId: input.ArticuloId, Cantidad: input.Cantidad}
	models.DB.Create(&cantidad)

	c.JSON(http.StatusOK, gin.H{"resultado": cantidad})
}
