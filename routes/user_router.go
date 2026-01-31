package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", handlers.GetUserHandler)
}
