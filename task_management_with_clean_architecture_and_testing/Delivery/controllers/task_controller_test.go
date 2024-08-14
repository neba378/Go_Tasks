package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"task_with_clean_arc_and_test/domain"
	"task_with_clean_arc_and_test/infrastructures"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockTaskUsecase struct {
	mock.Mock
}

func (m *MockTaskUsecase) GetTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) GetTaskByID(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) AddTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUsecase) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskUsecase) UpdateTask(id string, task domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

// Test suite for TaskHandler
type TaskHandlerTestSuite struct {
	suite.Suite
	router      *gin.Engine
	mockUsecase *MockTaskUsecase
	handler     *TaskHandler
}

func (suite *TaskHandlerTestSuite) SetupTest() {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	suite.router = gin.Default()

	// Create a mock usecase and handler
	suite.mockUsecase = new(MockTaskUsecase)
	suite.handler = &TaskHandler{usecase: suite.mockUsecase}

	allowed := suite.router.Group("")
	allowed.Use(infrastructures.AuthUser())
	allowed.GET("/tasks", suite.handler.GetTasks)
	allowed.GET("/tasks/:id", suite.handler.GetTaskByID)

	// Routes for admin users
	protected := suite.router.Group("/admin")
	protected.Use(infrastructures.AuthMiddleware("admin"))
	protected.PUT("/tasks/:id", suite.handler.UpdateTask)
	protected.DELETE("/tasks/:id", suite.handler.DeleteTask)
	protected.POST("/tasks", suite.handler.AddTask)
}

func (suite *TaskHandlerTestSuite) TestGetTasks_Success() {
	// Set a specific time for testing
	dueDate1 := time.Date(2024, 8, 13, 16, 29, 6, 0, time.Local)
	dueDate2 := time.Date(2024, 8, 13, 16, 29, 6, 0, time.Local)

	// Mock tasks
	tasks := []domain.Task{
		{
			ID:          "1",
			Title:       "Task 1",
			Description: "Description 1",
			DueDate:     dueDate1,
			Status:      "pending",
		},
		{
			ID:          "2",
			Title:       "Task 2",
			Description: "Description 2",
			DueDate:     dueDate2,
			Status:      "completed",
		},
	}

	// Mock usecase
	suite.mockUsecase.On("GetTasks").Return(tasks, nil)

	// Generate a valid JWT token for an authenticated user
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "test_user",
		Role:     "user", // Ensure this role has permissions for the endpoint
	}
	token, err := infrastructures.GenerateToken(user)
	suite.NoError(err)

	// Create a new GET request with the token
	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Print the raw JSON response for debugging
	fmt.Println("Response Body:", w.Body.String())

	// Unmarshal the response body into a slice of tasks
	var returnedTasks []domain.Task
	err = json.Unmarshal(w.Body.Bytes(), &returnedTasks)
	if err != nil {
		suite.Fail("Failed to unmarshal response: %v", err)
		return
	}

	// Assert the response and ensure it matches the expected values
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Len(suite.T(), returnedTasks, len(tasks))
	for i, task := range tasks {
		assert.Equal(suite.T(), task.ID, returnedTasks[i].ID)
		assert.Equal(suite.T(), task.Title, returnedTasks[i].Title)
		assert.Equal(suite.T(), task.Description, returnedTasks[i].Description)
		assert.Equal(suite.T(), task.Status, returnedTasks[i].Status)
		assert.True(suite.T(), task.DueDate.Equal(returnedTasks[i].DueDate))
	}
}

func (suite *TaskHandlerTestSuite) TestGetTaskByID_Success() {
	fixedTime := time.Date(2024, 8, 13, 16, 29, 6, 0, time.Local)
	task := domain.Task{
		ID:          "1",
		Title:       "Test Task",
		Description: "Task Description",
		DueDate:     fixedTime,
		Status:      "pending",
	}

	// Set up the mock to expect a call with the ID "1" and return the mock task
	suite.mockUsecase.On("GetTaskByID", "1").Return(task, nil)

	// Generate a valid JWT token for an authenticated user
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "test_user",
		Role:     "user", // Ensure this role has permissions for the endpoint
	}
	token, err := infrastructures.GenerateToken(user)
	suite.NoError(err)

	// Create a new GET request with the token
	req, err := http.NewRequest(http.MethodGet, "/tasks/1", nil)
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Assert the response status
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	// Prepare the expected response body
	expectedBody := `{"id":"1","title":"Test Task","description":"Task Description","due_date":"` +
		task.DueDate.Format(time.RFC3339) + `","status":"pending"}`

	// Compare the expected body with the actual response
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *TaskHandlerTestSuite) TestGetTaskByID_NotFound() {
	// Generate a valid JWT token for an authenticated user
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "test_user",
		Role:     "user", // Ensure this role has permissions for the endpoint
	}
	token, err := infrastructures.GenerateToken(user)
	suite.NoError(err)

	// Mock the usecase to return an error indicating the task was not found
	suite.mockUsecase.On("GetTaskByID", "1").Return(domain.Task{}, errors.New("Task not found"))

	// Create a new GET request with the token
	req, err := http.NewRequest(http.MethodGet, "/tasks/1", nil)
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Assert the response status
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)

	// Prepare the expected response body
	expectedBody := `{"error":"Task not found"}`

	// Compare the expected body with the actual response
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *TaskHandlerTestSuite) TestAddTask_Success() {
	// Create a new task to be added
	newTask := domain.Task{
		ID:          "1",
		Title:       "New Task",
		Description: "New Task Description",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	// Marshal the new task to JSON
	payload, _ := json.Marshal(newTask)

	// Mock the use case to expect the task addition and return no error
	suite.mockUsecase.On("AddTask", mock.MatchedBy(func(task domain.Task) bool {
		return task.ID == newTask.ID &&
			task.Title == newTask.Title &&
			task.Description == newTask.Description &&
			task.Status == newTask.Status
	})).Return(nil)
	// Generate a valid JWT token for an admin user
	adminUser := domain.User{
		ID:       primitive.NewObjectID(), // Generate a new ObjectID
		Username: "admin_user",
		Role:     "admin",
	}
	token, err := infrastructures.GenerateToken(adminUser)
	suite.NoError(err)
	fmt.Println("Generated Token: ", token) // Debugging the generated token

	// Create a new POST request with the token in the Authorization header
	req, _ := http.NewRequest("POST", "/admin/tasks", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()
	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Debugging the response body
	suite.T().Logf("Response Status Code: %d", w.Code)
	suite.T().Logf("Response Body: %s", w.Body.String())

	// Check that the status code is 201 Created
	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	// Define the expected response body
	expectedBody := `{"message":"Task created"}`

	// Check that the response body matches the expected JSON
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *TaskHandlerTestSuite) TestAddTask_BadRequest() {
	// Create an invalid or empty JSON payload
	invalidPayload := []byte(``)

	// Mock the usecase to ensure it doesn't get called
	suite.mockUsecase.On("AddTask", mock.Anything).Return(nil).Maybe()

	// Generate a valid JWT token for an authenticated user
	user := domain.User{
		ID:       primitive.NewObjectID(), // Generate a new ObjectID
		Username: "test_user",
		Role:     "admin", // Ensure this role has permissions for the endpoint
	}
	token, err := infrastructures.GenerateToken(user)
	suite.NoError(err)
	fmt.Println("Generated Token: ", token) // Debugging the generated token

	// Create a new POST request with the invalid payload and token
	req, err := http.NewRequest(http.MethodPost, "/admin/tasks", bytes.NewBuffer(invalidPayload))
	suite.NoError(err) // Ensure there is no error creating the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token) // Set the Authorization header

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Assert that the response status code is 400 Bad Request
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	// Assert that the response body contains the appropriate error message
	expectedBody := `{"error":"EOF"}`
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *TaskHandlerTestSuite) TestDeleteTask_Success() {
	suite.mockUsecase.On("DeleteTask", "1").Return(nil)
	user := domain.User{
		ID:       primitive.NewObjectID(), // Generate a new ObjectID
		Username: "test_user",
		Role:     "admin", // Ensure this role has permissions for the endpoint
	}
	token, err := infrastructures.GenerateToken(user)
	suite.NoError(err)
	fmt.Println("Generated Token: ", token) // Debugging the generated token

	// Create a new POST request with the task payload and token
	req, err := http.NewRequest(http.MethodDelete, "/admin/tasks/1", nil)
	suite.NoError(err) // Ensure there is no error creating the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	expectedBody := `{"message":"successfully deleted!"}`
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *TaskHandlerTestSuite) TestDeleteTask_InternalServerError() {
	// Mock the usecase to simulate an internal server error
	suite.mockUsecase.On("DeleteTask", "10").Return(errors.New("Task not found"))

	// Generate a valid JWT token for an authenticated user
	user := domain.User{
		ID:       primitive.NewObjectID(), // Generate a new ObjectID
		Username: "test_user",
		Role:     "admin", // Ensure this role has permissions for the endpoint
	}
	token, err := infrastructures.GenerateToken(user)
	suite.NoError(err)
	fmt.Println("Generated Token: ", token) // Debugging the generated token

	// Create a new DELETE request with the task ID and token
	req, err := http.NewRequest(http.MethodDelete, "/admin/tasks/10", nil)
	suite.NoError(err) // Ensure there is no error creating the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Print the raw response body for debugging
	fmt.Println("Response Body: ", w.Body.String())

	// Assert that the response status code is 500 Internal Server Error
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)

	// Assert that the response body contains the appropriate error message
	expectedBody := `{"error":"Task not found"}`
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *TaskHandlerTestSuite) TestUpdateTask_Success() {
	newTask := domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     time.Now(),
		Status:      "completed",
	}
	payload, _ := json.Marshal(newTask)
	suite.mockUsecase.On("UpdateTask", "1", mock.MatchedBy(func(task domain.Task) bool {
		return task.ID == newTask.ID &&
			task.Title == newTask.Title &&
			task.Description == newTask.Description &&
			task.Status == newTask.Status
	})).Return(nil)

	adminUser := domain.User{
		ID:       primitive.NewObjectID(), // Generate a new ObjectID
		Username: "admin_user",
		Role:     "admin",
	}
	token, err := infrastructures.GenerateToken(adminUser)
	suite.NoError(err)
	req, _ := http.NewRequest(http.MethodPut, "/admin/tasks/1", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	expectedBody := `{"message":"successfully updated!"}`
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *TaskHandlerTestSuite) TestUpdateTask_InternalServerError() {
	// Create a task to be updated
	updatedTask := domain.Task{
		ID:          "1",
		Title:       "Updated Task",
		Description: "Updated Task Description",
		DueDate:     time.Now(), // This will be generated dynamically
		Status:      "completed",
	}

	// Marshal the updated task to JSON
	payload, _ := json.Marshal(updatedTask)

	// Mock the usecase to simulate an internal server error
	suite.mockUsecase.On("UpdateTask", "1", mock.MatchedBy(func(task domain.Task) bool {
		// Match based on ID and other fields except DueDate
		return task.ID == updatedTask.ID &&
			task.Title == updatedTask.Title &&
			task.Description == updatedTask.Description &&
			task.Status == updatedTask.Status
	})).Return(errors.New("Task update failed"))

	// Generate a valid JWT token for an authenticated user
	user := domain.User{
		ID:       primitive.NewObjectID(), // Generate a new ObjectID
		Username: "test_user",
		Role:     "admin", // Ensure this role has permissions for the endpoint
	}
	token, err := infrastructures.GenerateToken(user)
	suite.NoError(err)
	fmt.Println("Generated Token: ", token) // Debugging the generated token

	// Create a new PUT request with the task payload and token
	req, err := http.NewRequest(http.MethodPut, "/admin/tasks/1", bytes.NewBuffer(payload))
	suite.NoError(err) // Ensure there is no error creating the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Print the raw response body for debugging
	fmt.Println("Response Body: ", w.Body.String())

	// Assert that the response status code is 500 Internal Server Error
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)

	// Assert that the response body contains the appropriate error message
	expectedBody := `{"error":"Task update failed"}`
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

// Main function to run the test suite
func TestTaskHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskHandlerTestSuite))
}
