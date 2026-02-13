package middlewares

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var logFile, _ = os.OpenFile("gin_log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

func LoggerFile() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		logLine := fmt.Sprintf("%s - [%s] \"%s %s\" %d %s %v\n",
			clientIP,
			start.Format(time.RFC1123),
			method,
			path,
			status,
			duration,
			"custom logging middleware ",
		)

		log.Println(logLine)
	}
}