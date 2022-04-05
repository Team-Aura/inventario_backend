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
		v1.GET("/user", controllers.ConseguirUsuarios)
		v1.GET("/user/:userID", controllers.ConseguirUsuario)
		v1.POST("/user", controllers.CrearUsuario)
		v1.PUT("/user/:userID", controllers.ActualizarUsuario)
		v1.DELETE("/user/:userID", controllers.BorrarUsuario)

		v1.POST("/login", controllers.Login)
	}

	app.Run()
}
