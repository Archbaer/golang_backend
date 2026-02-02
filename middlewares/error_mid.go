package middlewares

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(500, gin.H{
					"message": "Internal Server Error",
					"error":   err,
				})
			}
		}()
		c.Next()
	}
}
