package usecases_test

import (
	"testing"

	"task_with_clean_arc_and_test/domain"
	"task_with_clean_arc_and_test/infrastructures"
	"task_with_clean_arc_and_test/usecases"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockUserRepository is a mock implementation of the UserRepository interface.
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Register(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) LoginUser(username string) (domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(domain.User), args.Error(1)
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

func (m *MockUserRepository) Deactivate(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

func (m *MockUserRepository) UsernameExists(username string) (bool, error) {
	args := m.Called(username)
	return args.Bool(0), args.Error(1)
}

// UserUsecaseSuite defines the suite for UserUsecase tests.
type UserUsecaseSuite struct {
	suite.Suite
	mockRepo *MockUserRepository
	usecase  usecases.UserUsecase
}

// SetupTest sets up the test environment before each test in the suite.
func (suite *UserUsecaseSuite) SetupTest() {
	suite.mockRepo = new(MockUserRepository)
	suite.usecase = usecases.NewUserUsecase(suite.mockRepo)
}

// TestRegisterUser tests the Register method.
func (suite *UserUsecaseSuite) TestRegisterUser() {
	user := domain.User{Username: "testuser", Password: "password"}

	// Check if username already exists
	suite.mockRepo.On("UsernameExists", user.Username).Return(false, nil)

	// Hash the password before passing it to the mock
	hashedPassword, err := infrastructures.HashPassword(user.Password)
	suite.Require().Nil(err)

	user.Password = hashedPassword // Update the user with the hashed password

	// Use a wildcard matcher to ignore the specific value of the password
	suite.mockRepo.On("Register", mock.MatchedBy(func(u domain.User) bool {
		return u.Username == user.Username && u.Password != "" // Check username and ensure password is not empty
	})).Return(nil)

	err = suite.usecase.Register(user)

	suite.Assert().Nil(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestRegisterUserExists tests registration of an existing username.
func (suite *UserUsecaseSuite) TestRegisterUserExists() {
	user := domain.User{Username: "testuser", Password: "password"}

	suite.mockRepo.On("UsernameExists", user.Username).Return(true, nil)

	err := suite.usecase.Register(user)

	suite.Assert().EqualError(err, "username already exists")
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestLoginUserSuccess tests successful login.
func (suite *UserUsecaseSuite) TestLoginUserSuccess() {
	user := domain.User{Username: "testuser", Password: "password"}
	hashedPassword, _ := infrastructures.HashPassword(user.Password)

	// Set the expected user with the hashed password
	suite.mockRepo.On("LoginUser", user.Username).Return(domain.User{Username: user.Username, Password: hashedPassword}, nil)
	suite.mockRepo.On("UsernameExists", user.Username).Return(true, nil)

	// Attempt to login with the plain password
	token, err := suite.usecase.LoginUser(domain.User{Username: user.Username, Password: user.Password})

	suite.Assert().Nil(err)
	suite.Assert().NotEmpty(token)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestLoginUserInvalidPassword tests login with an invalid password.
func (suite *UserUsecaseSuite) TestLoginUserInvalidPassword() {
	user := domain.User{Username: "testuser", Password: "wrongpassword"}
	hashedPassword, _ := infrastructures.HashPassword("password") // Original password

	suite.mockRepo.On("UsernameExists", user.Username).Return(true, nil)
	suite.mockRepo.On("LoginUser", user.Username).Return(domain.User{Username: user.Username, Password: hashedPassword}, nil)

	err := infrastructures.CheckPasswordHash(user.Password, hashedPassword)
	suite.Assert().Error(err)

	_, err = suite.usecase.LoginUser(user)

	suite.Assert().EqualError(err, "invalid password")
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestLoginUserNotFound tests login for a non-existent user.
func (suite *UserUsecaseSuite) TestLoginUserNotFound() {
	user := domain.User{Username: "nonexistent", Password: "password"}

	suite.mockRepo.On("UsernameExists", user.Username).Return(false, nil)
	_, err := suite.usecase.LoginUser(user)

	suite.Assert().Error(err)
	suite.Contains(err.Error(), "user not found")
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestUserUsecaseSuite runs the test suite.
func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
