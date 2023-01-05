package controllers

import (
	"net/http"

	m "github.com/LordFeexz/Nusa-Tech/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func Register(c *gin.Context) {
	var user m.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user.Password = string(hashedPassword)

	m.Db.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"message": user})
}
