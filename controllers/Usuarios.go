package controllers

import (
	"TeamAura/go/ORM/models"
	"TeamAura/go/ORM/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ConseguirUsuarios(c *gin.Context) {
	var usuarios []models.Usuarios
	models.DB.Find(&usuarios)
	c.JSON(http.StatusOK, gin.H{"usuarios": usuarios})
}

func ConseguirUsuario(c *gin.Context) {
	var usuario models.Usuarios
	id := c.Param("userID")
	models.DB.Find(&usuario, id)
	c.JSON(http.StatusOK, gin.H{"usuario": usuario})

}

func CrearUsuario(c *gin.Context) {
	var inputUsuario UserInput
	if err := c.ShouldBindJSON(&inputUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var passswrd_hash, err = HashPassword(inputUsuario.Contrasenia)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No se ha podido crear el usuario"})
	} else {
		// Creamos el articulo
		usuario := models.Usuarios{NombreUsuario: inputUsuario.NombreUsuario, Contrasenia: passswrd_hash, Activo: true}
		models.DB.Create(&usuario)
		c.JSON(http.StatusOK, gin.H{"message": "Usuario creado con exito"})
	}

}

func ActualizarUsuario(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var usuario models.Usuarios

	id := c.Param("UserID")
	models.DB.Find(&usuario, id)
	var passswrd_hash, err = HashPassword(input.Contrasenia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No se ha podido actualizar el usuario"})
		return
	}
	usuario.NombreUsuario = input.NombreUsuario
	usuario.Contrasenia = passswrd_hash
	models.DB.Save(&usuario)
	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado con exito"})
}

func BorrarUsuario(c *gin.Context) {
	var usuario models.Usuarios
	id := c.Param("userID")

	models.DB.Delete(&usuario, id)
	c.JSON(http.StatusOK, gin.H{"message": "Usuario borrado con exito"})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *gin.Context) {
	var usuario models.Usuarios
	var loginVals UserInput
	if err := c.ShouldBind(&loginVals); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nombreUsuario := loginVals.NombreUsuario
	password := loginVals.Contrasenia

	models.DB.Where("nombre_usuario = ?", nombreUsuario).Find(&usuario)

	if CheckPasswordHash(password, usuario.Contrasenia) {
		var token = services.JWTAuthService().GenerateToken(usuario.NombreUsuario, true)
		var usuarioToken models.UsuarioToken
		usuarioToken.Token = token
		usuarioToken.Usuario = usuario
		models.DB.Create(&usuarioToken)

		c.JSON(http.StatusOK, gin.H{"message": token})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al iniciar sesion"})
	}

}
