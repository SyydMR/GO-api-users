package controllers

import (
	"net/http"
	"time"
    "github.com/golang-jwt/jwt/v5"

	"github.com/SyydMR/GO-api-users/database"
	"github.com/SyydMR/GO-api-users/models"
	"github.com/gin-gonic/gin"
)
var jwtKey = []byte("syydmr")

func GenerateToken(username string) (string, error) {
    claims := &jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := user.HashPassword(user.Password); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}
func Login(c *gin.Context) {
    var user models.User
    var loginUser models.User

    if err := c.ShouldBindJSON(&loginUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Where("email = ?", loginUser.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if err := user.CheckPassword(loginUser.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := GenerateToken(user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
func ProtectedEndpoint(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "You have access to this protected route!"})
}