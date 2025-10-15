package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/nacl/auth"
	"net/http"
	"shop.go/config"
	"shop.go/models"
)

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User

	if err := config.DB.Preload("Roles").Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "نام کاربری یا کلمه عبور اشتباه است"})
		return
	}

	if !CheckPasswordHash(req.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "نام کاربری یا کلمه عبور اشتباه است"})
		return
	}

	var roleNames []string
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}

	token, err := auth.GenerateJWT(user.ID, roleNames)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "خطای سرور "})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
