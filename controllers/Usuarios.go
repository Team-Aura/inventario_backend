package controllers

import (
	"TeamAura/go/ORM/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConseguirUsuarios(c *gin.Context) {
	var usuarios []models.Usuarios
	models.DB.Find(&usuarios)
	c.JSON(http.StatusOK, gin.H{"usuarios": usuarios})
}
