package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"task_with_clean_arc_and_test/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository interface {
	GetOne(id string) (domain.Task, error)
	GetAll() ([]domain.Task, error)
	Add(task domain.Task) error
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

func (r *taskRepository) GetOne(id string) (domain.Task, error) {
	filter := bson.D{{Key: "id", Value: id}}
	var res domain.Task
	err := r.collection.FindOne(context.TODO(), filter).Decode(&res)
	return res, err
}

func (r *taskRepository) GetAll() ([]domain.Task, error) {
	findOption := options.Find()
	var tasks []domain.Task
	curr, err := r.collection.Find(context.TODO(), bson.D{{}}, findOption) // the filter is not applied to get the whole task

	if err != nil {
		return nil, err
	}

	for curr.Next(context.TODO()) { //iterates till nothing is left
		var element domain.Task
		err := curr.Decode(&element)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, element)
	}
	return tasks, nil
}

func (r *taskRepository) Add(task domain.Task) error {
	// Retrieve all tasks and sort them by ID in descending order
	opts := options.Find().SetSort(bson.D{{Key: "id", Value: -1}})
	cursor, err := r.collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	// Initialize the LastID
	LastID := 0

	// Iterate through the tasks to find the highest ID
	for cursor.Next(context.TODO()) {
		var existingTask domain.Task
		if err := cursor.Decode(&existingTask); err != nil {
			return err
		}

		// Convert the existing ID to an integer
		id, err := strconv.Atoi(existingTask.ID)
		if err != nil {
			return err
		}

		// Update LastID if the current ID is higher
		if id > LastID {
			LastID = id
		}
	}

	// Increment the LastID to get the new ID
	LastID++
	task.ID = strconv.Itoa(LastID)
	task.Status = "Pending"
	task.DueDate = time.Now()
	if task.Title == "" || task.Description == "" {
		return errors.New("please provide a title and description")
	}
	_, err = r.collection.InsertOne(context.TODO(), task)
	return err
}

func (r *taskRepository) Delete(id string) error {
	result, err := r.collection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	if result.DeletedCount == 0 {
		return fmt.Errorf("task with id %s not found", id)
	}
	return err // deleted success
}

func (r *taskRepository) Update(id string, task domain.Task) error {
	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.M{"title": task.Title, "description": task.Description}}}
	if task.Title == "" || task.Description == "" {
		return errors.New("please provide a title and description")
	}
	result, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if result.MatchedCount == 0 {
		return fmt.Errorf("task with id %s not found", id)
	}

	return err // returns nill if the task is in there
}
