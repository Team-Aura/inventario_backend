package main

import (
	"TeamAura/go/ORM/controllers"
	"TeamAura/go/ORM/models"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	models.ConexionABBDD()
	//GETS
	app.GET("/usuarios", controllers.ConseguirUsuarios)
	app.GET("/articulos", controllers.ConseguirArticulos)
	app.GET("/almacenes", controllers.ConseguirAlmacenes)
	app.GET("/cantidades", controllers.ConseguirAlmacenesCantidades)
	//POSTS
	app.POST("/articulos/crear", controllers.CrearArticulos)
	app.POST("/almacenes/crear", controllers.CrearAlmacenes)
	app.POST("/cantidades/crear", controllers.CrearAlmacenAlmacenArticulosCantidades)
	app.Run()
}
