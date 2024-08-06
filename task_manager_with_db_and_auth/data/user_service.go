package data

import (
	"context"
	"errors"
	"fmt"

	"log"
	"os"
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
var collection *mongo.Collection

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
func Register(user models.User) (error){
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
	findOption := options.Find()
	cursor, err := collection.Find(context.TODO(),bson.D{{}},findOption)
	if err!=nil{
		return err
	}
	// fmt.Print(".,.,.,.,.,",err == mongo.ErrNoDocuments,",.,.,.,.,.")
	defer cursor.Close(context.TODO())
	if !cursor.Next(context.TODO()) {
		user.Role = "admin"
	}else {
		user.Role = "user"
	}
	
	user.Password = string(hashed)
	user.Activate = "true"
	insertOne, err := collection.InsertOne(context.TODO(),user)
	if err!=nil{
		return err
	}
	fmt.Print("\n inserted Id: ",insertOne.InsertedID,"\n")
	return nil
}

func RegisterAdmin(user models.User) (error){
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
	findOption := options.Find()
	_, err = collection.Find(context.TODO(),bson.D{{}},findOption)
	if err!=nil && err!=mongo.ErrNoDocuments{
		return err
	}
	user.Role = "admin"
	
	user.Password = string(hashed)
	user.Activate = "true"
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

func Activate(username string) error{
	filter := bson.D{{Key: "username", Value: username}}
	var res *models.User
	err := collection.FindOne(context.TODO(),filter).Decode(&res)
	if err != nil && err == mongo.ErrNoDocuments{
		return err
	} else if err!= nil{
		log.Fatal(err)
	}

	update := bson.D{{Key: "$set",Value: bson.M{"activate":"true"}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	// fmt.Println(res)
	if  err!=nil {
		log.Fatal(err)
	}
	return nil 
}

func DeActivate(username string) error{
	filter := bson.D{{Key: "username", Value: username}}
	var res *models.User
	err := collection.FindOne(context.TODO(),filter).Decode(&res)
	fmt.Print(res)
	if err != nil && err == mongo.ErrNoDocuments{
		return err
	} else if err!= nil{
		log.Fatal(err)
	}
	update := bson.D{{Key: "$set",Value: bson.M{"activate":"false"}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	// fmt.Println(res)
	if  err!=nil {
		log.Fatal(err)
	}
	return nil 
}

func UpdateUser(username string)(error){
	// fmt.Println(tasks)

	filter := bson.D{{Key: "username", Value: username}}

	var res *models.Task
	err := collection.FindOne(context.TODO(),filter).Decode(&res)
	if err != nil && err == mongo.ErrNoDocuments{
		return err
	} else if err!= nil{
		log.Fatal(err)
	}

	update := bson.D{{Key: "$set",Value: bson.M{"role":"admin"}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	// fmt.Println(res)
	if  err!=nil {
		log.Fatal(err)
	}
	fmt.Print(result)
	return nil // returns the updated tasks if the task is in there 	 
}