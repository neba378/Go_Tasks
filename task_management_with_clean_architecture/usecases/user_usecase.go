package usecases

import (
	"task_with_clean_arc/domain"
	"task_with_clean_arc/repository"
)

type UserUsecase interface {
	Register(user domain.User) error
    LoginUser(user domain.User) (string, error)
	RegisterAdmin(user domain.User) (error)
	UpdateUser(username string)(error)
	Activate(username string) error
	DeActivate(username string) error
}

type userUsecase struct{
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase{
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Register(user domain.User) error{
	return u.repo.Register(user)
}
func (u *userUsecase) LoginUser(user domain.User) (string, error){
	return u.repo.LoginUser(user)
}

func (u *userUsecase) RegisterAdmin(user domain.User) (error){
	return u.repo.RegisterAdmin(user)
}

func (u *userUsecase) UpdateUser(username string)(error){
	return u.repo.UpdateUser(username)
}

func (u *userUsecase) Activate(username string) error{
	return u.repo.Activate(username)
}

func (u *userUsecase) DeActivate(username string) error{
	return u.repo.DeActivate(username)
}