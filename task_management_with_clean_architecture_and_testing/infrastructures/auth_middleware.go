package infrastructures

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header missing!"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		fmt.Print(authParts)

		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userRole := claims["role"].(string)
			if userRole != requiredRole && requiredRole != "" {
				c.JSON(403, gin.H{"error": "Forbidden for you"})
				c.Abort()
				return
			}
			c.Set("user", claims) // store data in the ctx to make it accessible for the other handlers
		} else {
			c.JSON(401, gin.H{"error": "Invalid JWT claims"})
			c.Abort()
			return
		}

		c.Next() // used to proceed the request further
	}
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header missing!"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		fmt.Print(authParts)

		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			c.Set("user", claims) // store data in the ctx to make it accessible for the other handlers
		} else {
			c.JSON(401, gin.H{"error": "Invalid JWT claims"})
			c.Abort()
			return
		}
		c.Next() // used to proceed the request further
	}
}
