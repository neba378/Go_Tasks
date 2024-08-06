package repository

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"task_with_clean_arc/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Register(role string, user domain.User) error
	LoginUser(user domain.User) (string,error)
}

type userRepository struct{
	collection *mongo.Collection
	jwtSecret []byte
}

func NewUserRepository(client *mongo.Client) UserRepository{
	err := godotenv.Load(filepath.Join("configs", ".env"))
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
	return &userRepository{
		collection: client.Database("task_manager").Collection("users"),
		jwtSecret: []byte(os.Getenv("JWT_SECRET")),
	}

}

func (r *userRepository) Register(role string, user domain.User) error{
	got,err := r.usernameExists(user.Username)
	if got{
		return err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	
	if err!=nil{
		return err
	}

	opts := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})

	// Iterate ti get the last user
    var lastUser domain.User
    err = r.collection.FindOne(context.TODO(), bson.D{}, opts).Decode(&lastUser)
	if err != nil && err != mongo.ErrNoDocuments {
        log.Fatal("error, ",err)
    }
	
	// taking the last id from the db and converting to int.
	newID := "1"
	if lastUser.ID != "" {
        lastID, err := strconv.Atoi(lastUser.ID)
        if err != nil {
            log.Fatal(err)
        }
        newID = strconv.Itoa(lastID + 1)
    }
	
	user.Role = role
	user.ID = newID
	user.Password = string(hashed)

	_, err = r.collection.InsertOne(context.TODO(),user)
	return err
}

func (r *userRepository) LoginUser(user domain.User) (string,error){
	var existingUser domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		return "", err
	} else if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":       existingUser.ID,
		"username": existingUser.Username,
		"role": 	existingUser.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(r.jwtSecret)
	if err != nil {
		return "", err
	}

	return jwtToken ,nil
}


func (r *userRepository) usernameExists(username string) (bool, error) {
	var user domain.User
    err := r.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
    if err != nil && err != mongo.ErrNoDocuments {
		return false, err
    }
    return err == nil, nil
}
