package handlers

import (
	"github.com/gin-gonic/gin"
)

type UserData struct {
	Name     string
	Email    string
	Password int
}

var users = map[string]UserData{
	"user1": {
		Name:  "Alice",
		Email: "alice@bol.com.br",
	},
	"user2": {
		Name:  "Bob",
		Email: "bob.dd@gmail.com",
	},
}

func GetUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": users,
	})
}
