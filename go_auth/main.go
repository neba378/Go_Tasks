package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func main() {

	var userStore = make(map[string]*User)

	router := gin.Default()
	router.POST("/register", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err!=nil{
			ctx.JSON(400, gin.H{"error":"incorrect input"})
			return
		}

		bytePassword := []byte(user.Password)

		hashedUser, err := bcrypt.GenerateFromPassword(bytePassword,bcrypt.DefaultCost)
		if err!=nil{
			ctx.JSON(500, gin.H{"message":"internal error"})
			return

		}

		user.Password = string(hashedUser)
		userStore[user.Email] = &user
		ctx.JSON(200, gin.H{"message":"user successfully registered!", "user":user})

	})
	router.Run("localhost:8080")
}