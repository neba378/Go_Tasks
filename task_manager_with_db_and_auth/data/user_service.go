package data

import (
	"context"
	"errors"
	"fmt"

	"log"
	"os"
	"strconv"
	"task_manager_with_db_and_auth/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte

func init() {
	client := ConnectToDB()
	collection = UserCollection(client)
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

var collection *mongo.Collection

func UserCollection(client *mongo.Client) *mongo.Collection{
	return client.Database("task_manager").Collection("users") // creates a collection named tasks in task_manager db
}

func usernameExists(collection *mongo.Collection, username string) (bool, error) {
	var user models.User
    err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
    if err != nil && err != mongo.ErrNoDocuments {
		return false, err
    }
    return err == nil, nil
}
func Register(role string ,user models.User) (error){
	// var client = ConnectToDB()
	
	// collection := UserCollection(client)
	got,_ := usernameExists(collection, user.Username)
	if got{
		return errors.New("username exists")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	
	if err!=nil{
		return err
	}

	opts := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})

	// Iterate ti get the last user
    var lastUser models.User
    err = collection.FindOne(context.TODO(), bson.D{}, opts).Decode(&lastUser)
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
	fmt.Print("\n user password: ",user.Password)
	user.Password = string(hashed)

	insertOne, err := collection.InsertOne(context.TODO(),user)
	if err!=nil{
		return err
	}
	fmt.Print("\n inserted Id: ",insertOne.InsertedID,"\n")
	return nil
}

func CheckUser(user models.User) (string,error ){
	
	var existingUser models.User
	err := collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		return "", fmt.Errorf("username does not exist")
	} else if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	claims := jwt.MapClaims{
		"id":       existingUser.ID,
		"username": existingUser.Username,
		"role": 	existingUser.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	fmt.Println("message: ", "User successfully logged in, ", "token:", jwtToken)
	return jwtToken ,nil
}

