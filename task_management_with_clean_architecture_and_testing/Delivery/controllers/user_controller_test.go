package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"task_with_clean_arc_and_test/domain"
	"task_with_clean_arc_and_test/infrastructures"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Register(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUsecase) LoginUser(user domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockUserUsecase) RegisterAdmin(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUsecase) UpdateUser(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

func (m *MockUserUsecase) Activate(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

func (m *MockUserUsecase) Deactivate(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

// Test suite for UserHandler
type UserHandlerTestSuite struct {
	suite.Suite
	router      *gin.Engine
	mockUsecase *MockUserUsecase
	handler     *UserHandler
}

func (suite *UserHandlerTestSuite) SetupTest() {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	suite.router = gin.Default()

	// Create a mock usecase and handler
	suite.mockUsecase = new(MockUserUsecase)
	suite.handler = &UserHandler{Usecase: suite.mockUsecase}
	suite.router.POST("/register", suite.handler.RegisterUser)

	allowed := suite.router.Group("")
	allowed.POST("/login", suite.handler.LoginUser)

	// Routes for admin users
	protected := suite.router.Group("/admin")
	protected.Use(infrastructures.AuthMiddleware("admin"))
	protected.POST("/register", suite.handler.RegisterAdmin)
	protected.POST("/activate/:username", suite.handler.Activate)
	protected.POST("/deactivate/:username", suite.handler.DeActivate)
	protected.GET("/promote/:username", suite.handler.Promote)
}

func (suite *UserHandlerTestSuite) TestRegisterUser_Success() {
	user := domain.User{
		Username: "new_user",
		Password: "password123",
	}

	// Marshal the user to JSON
	payload, _ := json.Marshal(user)

	// Mock the use case to expect the user registration and return no error
	suite.mockUsecase.On("Register", mock.MatchedBy(func(u domain.User) bool {
		return u.Username == user.Username && u.Password == user.Password
	})).Return(nil)

	// Create a new POST request with the user data
	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(payload))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println("request:", req)
	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Check that the status code is 201 Created
	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	// Define the expected response body
	expectedBody := `{"message":"User registered"}`

	// Check that the response body matches the expected JSON
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *UserHandlerTestSuite) TestLoginUser_Success() {
	user := domain.User{
		Username: "test_user",
		Password: "password123",
	}
	token := "jwt_token"

	// Mock the use case to return a token
	suite.mockUsecase.On("LoginUser", user).Return(token, nil)

	// Create a new POST request with login credentials
	payload, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(payload))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)
	// Check that the status code is 200 OK
	// Check that the status code is 200 OK
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	// Define the expected response body
	expectedBody := `{"token":"` + token + `"}`

	// Check that the response body matches the expected JSON
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *UserHandlerTestSuite) TestRegisterAdmin_Success() {
	user := domain.User{
		Username: "admin_user",
		Password: "admin_password",
	}

	// Marshal the user to JSON
	payload, _ := json.Marshal(user)

	// Mock the use case to expect the admin registration and return no error
	suite.mockUsecase.On("RegisterAdmin", mock.MatchedBy(func(u domain.User) bool {
		return u.Username == user.Username && u.Password == user.Password
	})).Return(nil)

	// Generate a valid JWT token for an admin user
	adminUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "admin_user",
		Role:     "admin",
	}
	token, err := infrastructures.GenerateToken(adminUser)
	suite.NoError(err)

	// Create a new POST request with the admin registration data
	req, err := http.NewRequest(http.MethodPost, "/admin/register", bytes.NewBuffer(payload))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Check that the status code is 201 Created
	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	// Define the expected response body
	expectedBody := `{"message":"Successfully registered!"}`

	// Check that the response body matches the expected JSON
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *UserHandlerTestSuite) TestUpdateUser_Success() {

	user := domain.User{
		Username: "test_user",
		Password: "password123",
		Role:     "user",
	}
	username := "test_user"

	// Mock the use case to expect the update and return no error
	suite.mockUsecase.On("UpdateUser", username).Return(nil)

	// Generate a valid JWT token for an admin user
	adminUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "admin_user",
		Role:     "admin",
	}
	token, err := infrastructures.GenerateToken(adminUser)
	suite.NoError(err)
	payload, _ := json.Marshal(user)

	// Create a new POST request with the update data
	req, err := http.NewRequest(http.MethodGet, "/admin/promote/"+username, bytes.NewBuffer(payload))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Check that the status code is 200 OK
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	// Define the expected response body
	expectedBody := `{"message":"User updated"}`

	// Check that the response body matches the expected JSON
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *UserHandlerTestSuite) TestActivate_Success() {
	username := "test_user"
	user := domain.User{
		Username: "test_user",
		Password: "password123",
		Role:     "user",
	}
	payload, _ := json.Marshal(user)

	// Mock the use case to expect the activation and return no error
	suite.mockUsecase.On("Activate", username).Return(nil)

	// Generate a valid JWT token for an admin user
	adminUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "admin_user",
		Role:     "admin",
	}
	token, err := infrastructures.GenerateToken(adminUser)
	suite.NoError(err)

	// Create a new POST request with the activation data
	req, err := http.NewRequest(http.MethodPost, "/admin/activate/"+username, bytes.NewBuffer(payload))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Check that the status code is 200 OK
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	// Define the expected response body
	expectedBody := `{"message":"successfully activated!"}`

	// Check that the response body matches the expected JSON
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func (suite *UserHandlerTestSuite) TestDeactivate_Success() {
	username := "test_user"

	// Mock the use case to expect the deactivation and return no error
	suite.mockUsecase.On("Deactivate", username).Return(nil)

	// Generate a valid JWT token for an admin user
	adminUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "admin_user",
		Role:     "admin",
	}
	token, err := infrastructures.GenerateToken(adminUser)
	suite.NoError(err)

	// Create a new POST request with the deactivation data
	req, err := http.NewRequest(http.MethodPost, "/admin/deactivate/"+username, nil)
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	suite.router.ServeHTTP(w, req)

	// Check that the status code is 200 OK
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	// Define the expected response body
	expectedBody := `{"message":"successfully deactivated!"}`

	// Check that the response body matches the expected JSON
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
