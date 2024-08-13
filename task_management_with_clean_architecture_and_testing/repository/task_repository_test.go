package repository

import (
	"context"
	"testing"
	"time"

	"task_with_clean_arc_and_test/domain"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	client     *mongo.Client
	collection *mongo.Collection
	repo       TaskRepository
}

func (suite *TaskRepositoryTestSuite) SetupSuite() {
	// Setup MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	suite.NoError(err)
	suite.client = client
	suite.collection = client.Database("task_manager").Collection("tasks")
	suite.repo = NewTaskRepository(client)
}

func (suite *TaskRepositoryTestSuite) TearDownSuite() {
	// Drop the test database and disconnect
	err := suite.client.Database("task_manager").Drop(context.TODO())
	suite.NoError(err)
	err = suite.client.Disconnect(context.TODO())
	suite.NoError(err)
}

func (suite *TaskRepositoryTestSuite) SetupTest() {
	// Optionally clean up collection before each test
	_, err := suite.collection.DeleteMany(context.TODO(), bson.D{{}})
	suite.NoError(err)
}

func (suite *TaskRepositoryTestSuite) TestGetOne() {
	task := domain.Task{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"}
	_, err := suite.collection.InsertOne(context.TODO(), task)
	suite.NoError(err)

	result, err := suite.repo.GetOne("1")
	suite.NoError(err)
	suite.Equal(task.ID, result.ID)
	suite.Equal(task.Title, result.Title)
}

func (suite *TaskRepositoryTestSuite) TestGetAll() {
	tasks := []domain.Task{
		{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"},
		{ID: "2", Title: "Task 2", Description: "Description 2", DueDate: time.Now(), Status: "Completed"},
	}
	for _, task := range tasks {
		_, err := suite.collection.InsertOne(context.TODO(), task)
		suite.NoError(err)
	}

	results, err := suite.repo.GetAll()
	suite.NoError(err)
	suite.Len(results, 2)
}

func (suite *TaskRepositoryTestSuite) TestAdd() {
	task := domain.Task{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"}
	err := suite.repo.Add(task)
	suite.NoError(err)

	var result domain.Task
	err = suite.collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: task.ID}}).Decode(&result)
	suite.NoError(err)
	suite.Equal(task.ID, result.ID)
}

func (suite *TaskRepositoryTestSuite) TestDelete() {
	task := domain.Task{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"}
	_, err := suite.collection.InsertOne(context.TODO(), task)
	suite.NoError(err)

	err = suite.repo.Delete("1")
	suite.NoError(err)

	count, err := suite.collection.CountDocuments(context.TODO(), bson.D{{Key: "id", Value: "1"}})
	suite.NoError(err)
	suite.Equal(int64(0), count)
}

func (suite *TaskRepositoryTestSuite) TestUpdate() {
	task := domain.Task{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"}
	_, err := suite.collection.InsertOne(context.TODO(), task)
	suite.NoError(err)

	updatedTask := domain.Task{Title: "Updated Title", Description: "Updated Description"}
	err = suite.repo.Update("1", updatedTask)
	suite.NoError(err)

	var result domain.Task
	err = suite.collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: "1"}}).Decode(&result)
	suite.NoError(err)
	suite.Equal(updatedTask.Title, result.Title)
	suite.Equal(updatedTask.Description, result.Description)
}

func (suite *TaskRepositoryTestSuite) TestGetOne_NotFound() {
	result, err := suite.repo.GetOne("13")
	suite.Error(err)
	suite.Equal(domain.Task{}, result)
}

func (suite *TaskRepositoryTestSuite) TestGetOne_InvalidIDFormat() {
	result, err := suite.repo.GetOne("invalid_id_format")
	suite.Error(err)
	suite.Equal(domain.Task{}, result)
}

func (suite *TaskRepositoryTestSuite) TestGetAll_EmptyCollection() {
	// Ensure the collection is empty
	err := suite.collection.Drop(context.TODO())
	suite.NoError(err)

	result, err := suite.repo.GetAll()

	// Expect no error but an empty slice
	suite.NoError(err)
	suite.Equal(0, len(result))
}

func (suite *TaskRepositoryTestSuite) TestAdd_InvalidTaskData() {
	// Missing Title
	task := domain.Task{ID: "2", Description: "Description 1", DueDate: time.Now(), Status: "Pending"}
	err := suite.repo.Add(task)
	suite.Error(err)
}

func (suite *TaskRepositoryTestSuite) TestDelete_NotFound() {
	err := suite.repo.Delete("12000")
	suite.Error(err)
}

func (suite *TaskRepositoryTestSuite) TestDelete_InvalidIDFormat() {
	err := suite.repo.Delete("invalid_id_format")
	suite.Error(err)
}

func (suite *TaskRepositoryTestSuite) TestUpdate_NotFound() {
	task := domain.Task{Title: "Updated Title", Description: "Updated Description"}
	err := suite.repo.Update("12000", task)
	suite.Error(err)
}

func (suite *TaskRepositoryTestSuite) TestUpdate_InvalidTaskData() {
	// Assuming an ID exists
	task := domain.Task{Title: "", Description: "Updated Description"} // Missing Title
	err := suite.repo.Update("1", task)
	suite.Error(err)
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
