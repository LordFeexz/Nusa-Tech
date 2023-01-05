package controllers

import (
	"net/http"

	m "github.com/LordFeexz/Nusa-Tech/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadData(c *gin.Context) {
	var users []m.User

	m.Db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"Users": users})
}

func ReadDataById(c *gin.Context) {
	var user m.User

	id := c.Param("id")

	if err := m.Db.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"User": user})
}
