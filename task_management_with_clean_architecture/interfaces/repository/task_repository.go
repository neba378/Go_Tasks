package repository

import (
	"context"
	"strconv"
	"task_with_clean_arc/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository interface {
	GetOne(id string) (domain.Task,error)
	GetAll() ([]domain.Task, error)
	Add(task domain.Task)  error
	Delete(id string) error
	Update(id string, task domain.Task) error
}

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(client *mongo.Client) TaskRepository {
    return &taskRepository{
        collection: client.Database("task_manager").Collection("tasks"),
    }
}

func (r *taskRepository)GetOne(id string) (domain.Task,error){
	filter := bson.D{{Key: "id",Value: id}}
	var res domain.Task
	err := r.collection.FindOne(context.TODO(),filter).Decode(&res)
	return res, err
}

func (r *taskRepository) GetAll() ([]domain.Task, error){
	findOption := options.Find()
	var tasks []domain.Task
	curr, err := r.collection.Find(context.TODO(),bson.D{{}},findOption)// the filter is not applied to get the whole task
	
	if err!=nil{
		return nil ,err
	}

	for curr.Next(context.TODO()){ //iterates till nothing is left
		var element domain.Task
		err:=curr.Decode(&element)
		if err!=nil{
			return nil,err
		}
		tasks = append(tasks,element)
	}
    return tasks, nil
}

func (r *taskRepository) Add(task domain.Task)  error{
	opts := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})

	// Iterate ti get the last task
    var lastTask domain.Task
    err := r.collection.FindOne(context.TODO(), bson.D{}, opts).Decode(&lastTask)
	// fmt.Println("one||||||\n")
	if err != nil && err != mongo.ErrNoDocuments {
        return err
    }
	task.Status = "Pending"

	// taking the last id from the db and converting to int.
	LastID := 0
    if lastTask.ID != "" {
        LastID, err = strconv.Atoi(lastTask.ID)
        if err != nil {
            return err
        }
    }

	// add one to the last id to create new one
	LastID+=1
	t := strconv.Itoa(LastID)
	task.ID = t
	task.DueDate = time.Now()
	_, err = r.collection.InsertOne(context.TODO(),task)

	return err // returns nil after successful addition of new task

}

func (r *taskRepository) Delete(id string) error{
	_, err := r.collection.DeleteOne(context.TODO(), bson.D{{Key: "id",Value: id}})
	return err // deleted success
}

func (r *taskRepository) Update(id string, task domain.Task) error{
	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{{Key: "$set",Value: bson.M{"title":task.Title,"description":task.Description}}}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)

	return err // returns the updated tasks if the task is in there 
}
