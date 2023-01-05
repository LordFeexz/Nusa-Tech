package routes

import (
	c "github.com/LordFeexz/Nusa-Tech/controllers"
	m "github.com/LordFeexz/Nusa-Tech/models"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	m.Connect()

	r.GET("/users", c.ReadData)

	r.GET("/users/:id", c.ReadDataById)

	r.Run()
}
