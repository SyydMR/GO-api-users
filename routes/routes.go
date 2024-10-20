package routes

import (
	"github.com/SyydMR/GO-api-users/controllers"
	"github.com/SyydMR/GO-api-users/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.GET("/users", controllers.GetUsers)
    router.POST("/users", controllers.CreateUser)
    router.POST("/login", controllers.Login)
	router.GET("/protected", middlewares.AuthMiddleware(), controllers.ProtectedEndpoint)
}