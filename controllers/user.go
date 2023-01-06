package controllers

import (
	"net/http"

	h "github.com/LordFeexz/Nusa-Tech/helpers"
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

	email, password := c.PostForm("email"), c.PostForm("password")

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user.Email = email

	user.Password = string(hashedPassword)

	m.Db.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"message": user})
}

func Login(c *gin.Context) {
	var user m.User

	email, password := c.PostForm("email"), c.PostForm("password")

	m.Db.Where("email = ?", email).First(&user)

	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid email/password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid email/password"})
		return
	}

	access_token, _ := h.CreateToken(user)

	c.JSON(http.StatusOK, gin.H{"access_token": access_token})
}
