package usecases

import (
	"task_with_clean_arc_and_test/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockTaskRepository is a mock implementation of the TaskRepository interface.
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetAll() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetOne(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) Add(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskRepository) Update(id string, task domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

// TaskUsecaseSuite defines the suite for TaskUsecase tests.
type TaskUsecaseSuite struct {
	suite.Suite
	mockRepo *MockTaskRepository
	usecase  TaskUsecase
}

// SetupTest sets up the test environment before each test in the suite.
func (suite *TaskUsecaseSuite) SetupTest() {
	suite.mockRepo = new(MockTaskRepository)
	suite.usecase = NewTaskUsecase(suite.mockRepo)
}

// TestGetTasks tests the GetTasks method.
func (suite *TaskUsecaseSuite) TestGetTasks() {
	mockTasks := []domain.Task{
		{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"},
		{ID: "2", Title: "Task 2", Description: "Description 2", DueDate: time.Now(), Status: "Completed"},
	}
	suite.mockRepo.On("GetAll").Return(mockTasks, nil)

	tasks, err := suite.usecase.GetTasks()

	suite.Assert().Nil(err)
	suite.Assert().NotEmpty(tasks)
	suite.Assert().Equal(len(mockTasks), len(tasks))
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestGetTaskByID tests the GetTaskByID method.
func (suite *TaskUsecaseSuite) TestGetTaskByID() {
	task := domain.Task{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"}
	suite.mockRepo.On("GetOne", "1").Return(task, nil)

	returnedTask, err := suite.usecase.GetTaskByID("1")

	suite.Assert().Nil(err)
	suite.Assert().Equal(task, returnedTask)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestAddTask tests the AddTask method.
func (suite *TaskUsecaseSuite) TestAddTask() {
	task := domain.Task{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"}
	suite.mockRepo.On("Add", task).Return(nil)

	err := suite.usecase.AddTask(task)

	suite.Assert().Nil(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestDeleteTask tests the DeleteTask method.
func (suite *TaskUsecaseSuite) TestDeleteTask() {
	suite.mockRepo.On("Delete", "1").Return(nil)

	err := suite.usecase.DeleteTask("1")

	suite.Assert().Nil(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestUpdateTask tests the UpdateTask method.
func (suite *TaskUsecaseSuite) TestUpdateTask() {
	task := domain.Task{ID: "1", Title: "Updated Task", Description: "Updated Description", DueDate: time.Now(), Status: "Completed"}
	suite.mockRepo.On("Update", "1", task).Return(nil)

	err := suite.usecase.UpdateTask("1", task)

	suite.Assert().Nil(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestTaskUsecaseSuite runs the test suite.
func TestTaskUsecaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
}
