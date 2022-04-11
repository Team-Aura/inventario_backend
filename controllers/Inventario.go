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

func ConseguirArticulo(c *gin.Context) {
	var articulo models.Articulos
	id := c.Param("articuloID")
	models.DB.Find(&articulo, id)
	c.JSON(http.StatusOK, gin.H{"articulo": articulo})

}

func ConseguirAlmacenes(c *gin.Context) {
	var almacenes []models.Almacenes
	models.DB.Preload("Usuario").Find(&almacenes)
	c.JSON(http.StatusOK, gin.H{"almacenes": almacenes})
}

func ConseguirAlmacen(c *gin.Context) {
	var almacen models.Almacenes
	id := c.Param("almacenID")
	models.DB.Find(&almacen, id)
	c.JSON(http.StatusOK, gin.H{"almacen": almacen})

}

func ConseguirAlmacenesCantidades(c *gin.Context) {
	var cantidades []models.AlmacenArticulosCantidades

	models.DB.Preload("Almacen").Preload("Articulo").Preload("Almacen.Usuario").Find(&cantidades)
	c.JSON(http.StatusOK, gin.H{"AlmacenArticulosCantidades": cantidades})
}

func ConseguirAlmacenesCantidad(c *gin.Context) {
	var almacenartcant models.AlmacenArticulosCantidades
	id := c.Param("ID")
	models.DB.Find(&almacenartcant, id)

	c.JSON(http.StatusOK, gin.H{"almacen": almacenartcant})

}

//INSERTS DE LAS DIFERENTES TABLAS

func CrearArticulo(c *gin.Context) {
	// Validamos los datos de entrada
	var input ArticulosInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Creamos el articulo
	articulo := models.Articulos{Nombre: input.Nombre, Codigo: input.Codigo}
	models.DB.Create(&articulo)

	c.JSON(http.StatusOK, gin.H{"resultado": articulo})
}

func CrearAlmacen(c *gin.Context) {
	// Validamos los datos de entrada
	var input AlmacenesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var usuario models.Usuarios
	models.DB.Find(&usuario, input.Usuario)
	// Creamos el articulo
	almacen := models.Almacenes{Usuario: usuario, Concepto: input.Concepto, Activo: true}
	models.DB.Create(&almacen)
	c.JSON(http.StatusOK, gin.H{"almacen": almacen})

}

func CrearAlmacenAlmacenArticulosCantidad(c *gin.Context) {
	// Validamos los datos de entrada
	var input AlmacenArticulosCantidadesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Creamos el articulo
	var almacen models.Almacenes
	models.DB.Find(&almacen, input.AlmacenId)
	var articulo models.Articulos
	models.DB.Find(&articulo, input.ArticuloId)
	cantidad := models.AlmacenArticulosCantidades{Almacen: almacen, Articulo: articulo, Cantidad: input.Cantidad}
	models.DB.Create(&cantidad)

	c.JSON(http.StatusOK, gin.H{"resultado": cantidad})
}

//Borrados

func BorrarArticulo(c *gin.Context) {
	// Validamos los datos de entrada
	var articulo models.Articulos
	id := c.Param("articuloID")

	models.DB.Delete(&articulo, id)
	c.JSON(http.StatusOK, gin.H{"message": "Artículo borrado con exito"})

}

func BorrarAlmacen(c *gin.Context) {
	// Validamos los datos de entrada
	var almacen models.Almacenes
	id := c.Param("almacenID")

	models.DB.Delete(&almacen, id)
	c.JSON(http.StatusOK, gin.H{"message": "Almacén borrado con exito"})
}

func BorrarAlmacenArticulosCantidad(c *gin.Context) {
	// Validamos los datos de entrada
	var almacenArticuloCantidad models.AlmacenArticulosCantidades
	id := c.Param("Id")
	models.DB.Delete(&almacenArticuloCantidad, id)
	c.JSON(http.StatusOK, gin.H{"message": "Cantidad y articulos borrados con exito"})
}

//Update

func ActualizarArticulo(c *gin.Context) {
	// Validamos los datos de entrada

	var input ArticulosInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var articulo models.Articulos

	id := c.Param("articuloID")
	models.DB.Find(&articulo, id)
	articulo.Nombre = input.Nombre
	articulo.Codigo = input.Codigo
	models.DB.Save(&articulo)
	c.JSON(http.StatusOK, gin.H{"message": "Artículo actualizado con exito"})

}

func ActualizarAlmacen(c *gin.Context) {
	// Validamos los datos de entrada
	var input AlmacenesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var almacen models.Almacenes

	id := c.Param("almacenID")
	models.DB.Find(&almacen, id)
	almacen.Concepto = input.Concepto
	almacen.Activo = input.Activo
	//almacen.UsuarioId = input.Usuario
	models.DB.Save(&almacen)
	c.JSON(http.StatusOK, gin.H{"message": "Almacén actualizado con exito"})
}

func ActualizarAlmacenArticulosCantidad(c *gin.Context) {
	// Validamos los datos de entrada
	var input AlmacenArticulosCantidadesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var almacenart models.AlmacenArticulosCantidades

	id := c.Param("Id")
	models.DB.Find(&almacenart, id)
	var almacen models.Almacenes
	models.DB.Find(&almacen, input.AlmacenId)
	var articulo models.Articulos
	models.DB.Find(&articulo, input.ArticuloId)
	//almacenart.AlmacenesID = uint(input.AlmacenId)
	//almacenart.ArticulosID = uint(input.ArticuloId)
	almacenart.Cantidad = input.Cantidad
	//models.DB.Save(&almacen)
	c.JSON(http.StatusOK, gin.H{"message": "Deposito actualizado con exito"})
}
