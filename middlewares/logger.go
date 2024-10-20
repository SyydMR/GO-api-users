package middlewares

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()

        c.Next()

        log.Printf("Request processed in %v", time.Since(startTime))
    }
}