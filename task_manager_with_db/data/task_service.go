package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"task_manager_with_db/models"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// this function is used to create a connection to mongodb:atlas by finding the uri from the .env file
func ConnectToDB() (*mongo.Client){

	err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }
	mongoURI := os.Getenv("MONGO_URI") //loading the mongo uri from the .env file
	// mongoURI := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(),clientOptions)

	if err!=nil{
		log.Fatal("Error during connecting to mongodb: ",err)
	}

	err = client.Ping(context.TODO(),nil) // Testing the connection

	if err!=nil{
		log.Fatal("Error during connecting to mongodb: ",err)
	}
	fmt.Print("successfully connected!")
	return client //sends client that is the connected db location
}

func TaskCollection(client *mongo.Client) *mongo.Collection{
	return client.Database("task_manager").Collection("tasks") // creates a collection named tasks in task_manager db
}

// this function is used to provide one task with given id
func GetOne(id string) (models.Task,error){
	var client = ConnectToDB()
	var collection = TaskCollection(client)
	var res *models.Task
	filter := bson.D{{Key: "id",Value: id}}
	err := collection.FindOne(context.TODO(),filter).Decode(&res)
	if err!= nil {
		return *res, err
	}
	return *res, nil
}
// this function is used to provide the whole tasks
func GetAllTasks() []models.Task {
	var client = ConnectToDB()
	var collection = TaskCollection(client)
	findOption := options.Find()
	var result []models.Task
	curr, err := collection.Find(context.TODO(),bson.D{{}},findOption)// the filter is not applied to get the whole task
	
	if err!=nil{
		log.Fatal("error searching for the tasks. ",err)
	}

	for curr.Next(context.TODO()){ //iterates till nothing is left
		var element models.Task
		err:=curr.Decode(&element)
		if err!=nil{
			log.Fatal("error while retrieving the tasks! ",err)
		}
		result = append(result,element)
	}
    return result
}

// this function is used in the adding of new task
func AddTask(newTask models.Task) error{
	var client = ConnectToDB()
	var collection = TaskCollection(client)
	opts := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})

	// Iterate ti get the last task
    var lastTask models.Task
    err := collection.FindOne(context.TODO(), bson.D{}, opts).Decode(&lastTask)
	// fmt.Println("one||||||\n")
	if err != nil && err != mongo.ErrNoDocuments {
        log.Fatal("error, ",err)
    }
	newTask.Status = "Pending"

	// taking the last id from the db and converting to int.
	LastID, err := strconv.Atoi(lastTask.ID)
    if err!= nil && err == mongo.ErrNoDocuments {
		LastID = 0
    } else if err!= nil{
		log.Fatal(err)
	}

	// add one to the last id to create new one
	LastID+=1
	t := strconv.Itoa(LastID)
	newTask.ID = t
	newTask.DueDate = time.Now()
	insertOne, err := collection.InsertOne(context.TODO(),newTask)
	if err!=nil{
		log.Fatal("Error during adding!")
	}
	fmt.Print("\n inserted Id: ",insertOne.InsertedID,"\n")
	return nil // returns tasks after successful addition of new task
	
}


// this function is used to remove a task with the given id
func DeleteTask(id string) (error) {
	var client = ConnectToDB()
	var collection = TaskCollection(client)
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{Key: "id",Value: id}})
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %v documents in the tasks collection\n", deleteResult.DeletedCount)

	return nil // deleted success
}

// this function is used to update task's title and description with the given id
func UpdateTask(id string,UpdatedTask models.Task)(error){
	// fmt.Println(tasks)
	var client = ConnectToDB()
	var collection = TaskCollection(client)

	filter := bson.D{{Key: "id", Value: id}}

	var res *models.Task
	err := collection.FindOne(context.TODO(),filter).Decode(&res)
	if err != nil && err == mongo.ErrNoDocuments{
		
		return err
	} else if err!= nil{
		log.Fatal(err)
	}

	update := bson.D{{Key: "$set",Value: bson.M{"title":UpdatedTask.Title,"description":UpdatedTask.Description}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	// fmt.Println(res)
	if  err!=nil {
		log.Fatal(err)
	}
	fmt.Print(result)
	return nil // returns the updated tasks if the task is in there 	 
}