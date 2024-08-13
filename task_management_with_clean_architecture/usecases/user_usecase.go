package usecases

import (
	"errors"
	"task_with_clean_arc/domain"
	"task_with_clean_arc/infrastructures"
	"task_with_clean_arc/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserUsecase interface {
	Register(user domain.User) error
	LoginUser(user domain.User) (string, error)
	RegisterAdmin(user domain.User) error
	UpdateUser(username string) error
	Activate(username string) error
	Deactivate(username string) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Register(user domain.User) error {
	// Check if username already exists
	exists, err := u.repo.UsernameExists(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already exists")
	}

	// Hash the user's password
	hashedPassword, err := infrastructures.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Set user role and activation status
	user.Activate = "true"
	err = u.repo.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) LoginUser(user domain.User) (string, error) {
	// Fetch the user from the repository by username
	existingUser, err := u.repo.LoginUser(user.Username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("user not found")
		}
		return "", err
	}

	// Check if the provided password matches the stored hashed password
	err = infrastructures.CheckPasswordHash(user.Password, existingUser.Password)
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Generate a JWT token for the authenticated user
	token, err := infrastructures.GenerateToken(existingUser)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userUsecase) RegisterAdmin(user domain.User) error {
	// Check if username already exists
	exists, err := u.repo.UsernameExists(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already exists")
	}

	// Hash the user's password
	hashedPassword, err := infrastructures.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	user.Role = "admin"
	user.Activate = "true"

	return u.repo.RegisterAdmin(user)
}

func (u *userUsecase) UpdateUser(username string) error {
	return u.repo.UpdateUser(username)
}

func (u *userUsecase) Activate(username string) error {
	return u.repo.Activate(username)
}

func (u *userUsecase) Deactivate(username string) error {
	return u.repo.Deactivate(username)
}
