package repository

import (
	"context"
	"errors"
	"log"
	"os"
	"task_with_clean_arc/domain"
	"task_with_clean_arc/infrastructures"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Register(user domain.User) error
	LoginUser(user domain.User) (string,error)
	RegisterAdmin(user domain.User) (error)
	UpdateUser(username string)(error)
	Activate(username string) error
	DeActivate(username string) error
}

type userRepository struct{
	collection *mongo.Collection
	jwtSecret []byte
}

func NewUserRepository(client *mongo.Client) UserRepository{
	err := godotenv.Load()
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
	return &userRepository{
		collection: client.Database("task_manager").Collection("users"),
		jwtSecret: []byte(os.Getenv("JWT_SECRET")),
	}

}

func (r *userRepository) Register(user domain.User) error{
	got,_ := r.usernameExists(user.Username)
	if got{
		return errors.New("username exists")
	}
	hashed, err := infrastructures.HashPassword(user.Password)
	
	if err!=nil{
		return err
	}
	findOption := options.Find()
	cursor, err := r.collection.Find(context.TODO(),bson.D{{}},findOption)
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
	_, err = r.collection.InsertOne(context.TODO(),user)
	if err!=nil{
		return err
	}
	// fmt.Print("\n inserted Id: ",insertOne.InsertedID,"\n")
	return nil
}

func (r *userRepository) LoginUser(user domain.User) (string,error){
	var existingUser domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		return "", err
	} else if err != nil {
		return "", err
	}

	err = infrastructures.CheckPasswordHash(user.Password, existingUser.Password)
	if err != nil {
		return "", err
	}
	
	jwtToken, err := infrastructures.GenerateToken(existingUser)
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
func (r *userRepository) RegisterAdmin(user domain.User) error {
	got, err := r.usernameExists(user.Username)
	if got {
		return err
	}
	hashed, err := infrastructures.HashPassword(user.Password)
	if err != nil {
		return err
	}
	findOption := options.Find()
	_, err = r.collection.Find(context.TODO(), bson.D{{}}, findOption)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	user.Role = "admin"

	user.Password = string(hashed)
	user.Activate = "true"
	_, err = r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	// fmt.Print("\n inserted Id: ", insertOne.InsertedID, "\n")
	return nil
}

func (r *userRepository) Activate(username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	var res *domain.User
	err := r.collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil && err == mongo.ErrNoDocuments {
		return err
	} else if err != nil {
		log.Fatal(err)
	}

	update := bson.D{{Key: "$set", Value: bson.M{"activate": "true"}}}

	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (r *userRepository) DeActivate(username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	var res *domain.User
	err := r.collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil && err == mongo.ErrNoDocuments {
		return err
	} else if err != nil {
		log.Fatal(err)
	}
	update := bson.D{{Key: "$set", Value: bson.M{"activate": "false"}}}

	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (r *userRepository) UpdateUser(username string) error {
	filter := bson.D{{Key: "username", Value: username}}

	var res *domain.Task
	err := r.collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil && err == mongo.ErrNoDocuments {
		return err
	} else if err != nil {
		log.Fatal(err)
	}

	update := bson.D{{Key: "$set", Value: bson.M{"role": "admin"}}}

	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Print(result)
	return nil // returns the updated tasks if the task is in there
}