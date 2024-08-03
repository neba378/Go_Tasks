package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID uint			`json:"id"`
	Email string 	`json:"email"`
	Password string `json:"password"`
}

func main() {

	var userStore = make(map[string]*User)
	jwtSecrete := []byte("h2ssh") // random secrete
	lastID := 1
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
		user.ID = uint(lastID)
		lastID++
		user.Password = string(hashedUser)
		userStore[user.Email] = &user
		ctx.JSON(200, gin.H{"message":"user successfully registered!", "user":user})

	})
	router.POST("/login",func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err!=nil{
			ctx.JSON(400, gin.H{"error":"incorrect input"})
			return
		}

		existingUser, ok := userStore[user.Email]
		// fmt.Print(bcrypt.CompareHashAndPassword([]byte(existingUser.Password),[]byte(user.Password)))
		if !ok || bcrypt.CompareHashAndPassword([]byte(existingUser.Password),[]byte(user.Password)) != nil{
			ctx.JSON(400, gin.H{"error":"incorrect email or password"})
			return
		}

		claim := jwt.MapClaims{
			"user_id": existingUser.ID,
			"email":   existingUser.Email,
		}

		//generate token
		token:= jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
		jwtToken,err := token.SignedString(jwtSecrete)

		if err!=nil{
			ctx.JSON(500, gin.H{"message":"internal error what"})
			return

		}

		ctx.JSON(200, gin.H{"message":"User successfully logged in", "token":jwtToken})

	})
	router.Run("localhost:8080")
}