package usecases

import (
	"errors"
	"task_with_clean_arc_and_test/domain"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Register(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) LoginUser(user domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) RegisterAdmin(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

func (m *MockUserRepository) Activate(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

func (m *MockUserRepository) DeActivate(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

type UserUsecaseSuite struct {
	suite.Suite
	mockRepo *MockUserRepository
	usecase  UserUsecase
}

func (suite *UserUsecaseSuite) SetupTest() {
	suite.mockRepo = new(MockUserRepository)
	suite.usecase = NewUserUsecase(suite.mockRepo)
}

func (suite *UserUsecaseSuite) TestRegister() {
	// Define a test user
	user := domain.User{Username: "testuser", Password: "password123"}

	// Mock the behavior of Register to return nil (no error)
	suite.mockRepo.On("Register", user).Return(nil)

	// Call the Register method of the usecase
	err := suite.usecase.Register(user)

	// Assert that no error occurred
	suite.NoError(err)

	// Ensure that the mock expectations were met
	suite.mockRepo.AssertExpectations(suite.T())
}

// Test registration with empty user data
func (suite *UserUsecaseSuite) TestRegisterEmptyUser() {
	user := domain.User{Username: "", Password: ""}
	suite.mockRepo.On("Register", user).Return(errors.New("invalid user"))

	err := suite.usecase.Register(user)
	suite.Error(err)
	suite.Contains(err.Error(), "invalid user")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestLoginUser() {
	// Define a test user and expected token
	user := domain.User{Username: "testuser", Password: "password123"}
	expectedToken := "mockToken123"

	// Mock the behavior of LoginUser to return the expected token and nil error
	suite.mockRepo.On("LoginUser", user).Return(expectedToken, nil)

	// Call the LoginUser method of the usecase
	token, err := suite.usecase.LoginUser(user)

	// Assert that the token matches the expected token and no error occurred
	suite.NoError(err)
	suite.Equal(expectedToken, token)

	// Ensure that the mock expectations were met
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestLoginUserWrongUser() {
	user := domain.User{Username: "wrong", Password: "wrongPassword"}
	suite.mockRepo.On("LoginUser", user).Return("", errors.New("invalid credentials"))
	token, err := suite.usecase.LoginUser(user)

	suite.Error(err)
	suite.Contains(err.Error(), "invalid credentials")
	suite.Empty(token)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestRegisterAdmin() {
	user := domain.User{Username: "adminuser", Password: "adminpassword"}
	suite.mockRepo.On("RegisterAdmin", user).Return(nil)

	err := suite.usecase.RegisterAdmin(user)
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestUpdateUser() {
	username := "testuser"
	suite.mockRepo.On("UpdateUser", username).Return(nil)

	err := suite.usecase.UpdateUser(username)
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestUpdateUserInvalid() {
	username := "nonexistentuser"
	suite.mockRepo.On("UpdateUser", username).Return(errors.New("user not found"))

	err := suite.usecase.UpdateUser(username)
	suite.Error(err)
	suite.Contains(err.Error(), "user not found")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestActivateUser() {
	username := "testuser"
	suite.mockRepo.On("Activate", username).Return(nil)

	err := suite.usecase.Activate(username)
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestDeActivateUser() {
	username := "testuser"
	suite.mockRepo.On("DeActivate", username).Return(nil)

	err := suite.usecase.DeActivate(username)
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestActivateUserInvalid() {
	username := "nonexistentuser"
	suite.mockRepo.On("Activate", username).Return(errors.New("user not found"))

	err := suite.usecase.Activate(username)
	suite.Error(err)
	suite.Contains(err.Error(), "user not found")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestDeActivateUserInvalid() {
	username := "nonexistentuser"
	suite.mockRepo.On("DeActivate", username).Return(errors.New("user not found"))

	err := suite.usecase.DeActivate(username)
	suite.Error(err)
	suite.Contains(err.Error(), "user not found")
	suite.mockRepo.AssertExpectations(suite.T())
}

// Run the test suite
func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
