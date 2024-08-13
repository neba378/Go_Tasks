package repository

import (
	"context"
	"task_with_clean_arc_and_test/domain"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (suite *UserRepositoryTestSuite) TestRegister_ExistingUser() {
	// Insert a user to test with
	existingUser := domain.User{Username: "existingUser", Password: "password123"}
	_, err := suite.collection.InsertOne(context.TODO(), existingUser)
	suite.NoError(err)

	// Attempt to register the same user
	err = suite.repo.Register(existingUser)
	suite.EqualError(err, "username exists")
}

func (suite *UserRepositoryTestSuite) TestLoginUser_InvalidCredentials() {
	// Attempt to login with an invalid username
	_, err := suite.repo.LoginUser(domain.User{Username: "invalidUser", Password: "wrongpassword"})
	suite.EqualError(err, "mongo: no documents in result")

	// Attempt to login with an invalid password
	// First register a valid user
	validUser := domain.User{Username: "validUser", Password: "password123"}
	err = suite.repo.Register(validUser)
	suite.NoError(err)

	// Attempt login with a wrong password
	_, err = suite.repo.LoginUser(domain.User{Username: "validUser", Password: "wrongpassword"})
	suite.EqualError(err, "crypto/bcrypt: hashedPassword is not the hash of the given password")
}

func (suite *UserRepositoryTestSuite) TestActivate_UserNotFound() {
	// Attempt to activate a non-existent user
	err := suite.repo.Activate("nonexistentUser")
	suite.EqualError(err, "mongo: no documents in result")
}

func (suite *UserRepositoryTestSuite) TestUpdateUser_UserNotFound() {
	// Attempt to update a non-existent user
	err := suite.repo.UpdateUser("nonexistentUser")
	suite.EqualError(err, "mongo: no documents in result")
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	// Cleanup after each test
	suite.collection.DeleteMany(context.TODO(), bson.D{{}})
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
