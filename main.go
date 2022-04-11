package main

import (
	"TeamAura/go/ORM/controllers"
	"TeamAura/go/ORM/models"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	models.ConexionABBDD()

	// Simple group: v1
	v1 := app.Group("/v1")
	{
		//Usuario
		v1.POST("/login", controllers.Login)
		v1.GET("/user", controllers.ConseguirUsuarios)
		v1.GET("/user/:userID", controllers.ConseguirUsuario)
		v1.POST("/user", controllers.CrearUsuario)
		v1.PUT("/user/:userID", controllers.ActualizarUsuario)
		v1.DELETE("/user/:userID", controllers.BorrarUsuario)
		//Articulos
		v1.GET("/articulo", controllers.ConseguirArticulos)
		v1.GET("/articulo/:articuloID", controllers.ConseguirArticulo)
		v1.POST("/articulo", controllers.CrearArticulo)
		v1.PUT("/articulo/:articuloID", controllers.ActualizarArticulo)
		v1.DELETE("/articulo/:articuloID", controllers.BorrarArticulo)
		//Almacenes
		v1.GET("/almacen", controllers.ConseguirAlmacenes)
		v1.GET("/almacen/:almacenID", controllers.ConseguirAlmacen)
		v1.POST("/almacen", controllers.CrearAlmacen)
		v1.PUT("/almacen/:almacenID", controllers.ActualizarAlmacen)
		v1.DELETE("/almacen/:almacenID", controllers.BorrarAlmacen)
		//Articulos Cantidades
		v1.GET("/cantidadarticulo", controllers.ConseguirAlmacenesCantidades)
		v1.GET("/cantidadarticulo/:ID", controllers.ConseguirAlmacenesCantidad)
		v1.POST("/cantidadarticulo", controllers.CrearAlmacenAlmacenArticulosCantidad)
		v1.PUT("/cantidadarticulo/:ID", controllers.ActualizarAlmacenArticulosCantidad)
		v1.DELETE("/cantidadarticulo/:ID", controllers.BorrarAlmacenArticulosCantidad)

	}

	app.Run()
}
