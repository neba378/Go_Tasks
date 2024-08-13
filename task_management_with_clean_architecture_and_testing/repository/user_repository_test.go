package repository

import (
	"context"
	"task_with_clean_arc_and_test/domain"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	client     *mongo.Client
	collection *mongo.Collection
	repo       UserRepository
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	suite.NoError(err)
	suite.client = client
	suite.collection = client.Database("task_manager").Collection("users")
	suite.repo = NewUserRepository(client)
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	err := suite.client.Disconnect(context.TODO())
	suite.NoError(err)
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	// Cleanup after each test
	suite.collection.DeleteMany(context.TODO(), bson.D{{}})
}

func (suite *UserRepositoryTestSuite) TestRegister_ExistingUser() {
	// Hash password for the existing user
	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	suite.NoError(err)

	existingUser := domain.User{Username: "existingUser", Password: string(hashedPassword)}
	_, err = suite.collection.InsertOne(context.TODO(), existingUser)
	suite.NoError(err)

	// Attempt to register the same user
	err = suite.repo.Register(existingUser)
	suite.EqualError(err, "username exists")
}

func (suite *UserRepositoryTestSuite) TestLoginUser_InvalidCredentials() {
	// Attempt to login with an invalid username
	_, err := suite.repo.LoginUser("invalidUser")
	suite.EqualError(err, "mongo: no documents in result")

	// Register a valid user
	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	suite.NoError(err)
	validUser := domain.User{Username: "validUser", Password: string(hashedPassword)}
	err = suite.repo.Register(validUser)
	suite.NoError(err)

	// Attempt login with a wrong password
	_, err = suite.repo.LoginUser("validUser")
	suite.NoError(err)
}

func (suite *UserRepositoryTestSuite) TestActivate_UserNotFound() {
	// Attempt to activate a non-existent user
	err := suite.repo.Activate("nonexistentUser")
	suite.EqualError(err, "user does not exist")
}

func (suite *UserRepositoryTestSuite) TestUpdateUser_UserNotFound() {
	// Attempt to update a non-existent user
	err := suite.repo.UpdateUser("nonexistentUser")
	suite.EqualError(err, "user does not exist")
}

func (suite *UserRepositoryTestSuite) TestUsernameExists() {
	// Register a user
	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	suite.NoError(err)
	user := domain.User{Username: "testUser", Password: string(hashedPassword)}
	err = suite.repo.Register(user)
	suite.NoError(err)

	// Check if username exists
	exists, err := suite.repo.UsernameExists("testUser")
	suite.NoError(err)
	suite.True(exists)

	// Check a non-existent username
	exists, err = suite.repo.UsernameExists("nonExistentUser")
	suite.NoError(err)
	suite.False(exists)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
