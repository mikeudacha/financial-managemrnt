package handlers

import (
	"financial-management-app/backend/internal/models"
	"financial-management-app/backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterUser(c *gin.Context) {
	var credentials Credentials
	if err := c.ShouldBind(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.Users{
		Username: cred.Username,
		Password: cred.Password,
	}
	services.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}

func LoginUser(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBind(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateToken(creds.Username, creds.Password)
	if err != nil {
		c.JSON
	}
}
