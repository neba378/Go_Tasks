package usecases

import (
	"task_with_clean_arc/domain"
	"task_with_clean_arc/interfaces/repository"
)

type UserUsecase interface {
	Register(role string, user domain.User) error
    LoginUser(user domain.User) (string, error)
}

type userUsecase struct{
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase{
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Register(role string, user domain.User) error{
	return u.repo.Register(role,user)
}
func (u *userUsecase) LoginUser(user domain.User) (string, error){
	return u.repo.LoginUser(user)
}