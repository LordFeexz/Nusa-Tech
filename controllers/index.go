package controllers

import (
	"net/http"

	m "github.com/LordFeexz/Nusa-Tech/models"
	"github.com/gin-gonic/gin"
)

func ReadData(c *gin.Context) {
	var users []m.User

	m.Db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"Users": users})
}
