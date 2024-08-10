package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskCreation(t *testing.T) {
	// Arrange
	id := "12345"
	title := "Sample Task"
	description := "This is a sample task description."
	dueDate := time.Now()
	status := "Pending"

	// Act
	task := Task{
		ID:          id,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      status,
	}

	// Assert
	assert.Equal(t, id, task.ID)
	assert.Equal(t, title, task.Title)
	assert.Equal(t, description, task.Description)
	assert.Equal(t, dueDate, task.DueDate)
	assert.Equal(t, status, task.Status)
}

func TestUserCreation(t *testing.T) {
	// Arrange
	id := primitive.NewObjectID()
	username := "testuser"
	password := "password123"
	role := "admin"
	activate := "yes"

	// Act
	user := User{
		ID:       id,
		Username: username,
		Password: password,
		Role:     role,
		Activate: activate,
	}

	// Assert
	assert.Equal(t, id, user.ID)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, password, user.Password)
	assert.Equal(t, role, user.Role)
	assert.Equal(t, activate, user.Activate)
}
