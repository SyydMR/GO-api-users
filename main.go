package main

import (
    "log"
    "github.com/SyydMR/GO-api-users/database"
    "github.com/SyydMR/GO-api-users/routes"
	"github.com/SyydMR/GO-api-users/middlewares"

    "github.com/gin-gonic/gin"
)

func main() {
    err := database.ConnectDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    r := gin.Default()

	r.Use(middlewares.LoggerMiddleware())
    routes.SetupRoutes(r)

    r.Run(":3000")
}