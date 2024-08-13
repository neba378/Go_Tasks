package repository

import (
	"context"
	"task_with_clean_arc/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Register(user domain.User) error
	LoginUser(username string) (domain.User, error)
	RegisterAdmin(user domain.User) error
	UpdateUser(username string) error
	Activate(username string) error
	Deactivate(username string) error
	UsernameExists(username string) (bool, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) UserRepository {
	return &userRepository{
		collection: client.Database("task_manager").Collection("users"),
	}
}

func (r *userRepository) Register(user domain.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *userRepository) RegisterAdmin(user domain.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *userRepository) Activate(username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.M{"activate": "true"}}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *userRepository) Deactivate(username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.M{"activate": "false"}}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *userRepository) UpdateUser(username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.M{"role": "admin"}}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *userRepository) UsernameExists(username string) (bool, error) {
	var user domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		return false, err
	}
	return err == nil, nil
}

func (r *userRepository) LoginUser(username string) (domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}