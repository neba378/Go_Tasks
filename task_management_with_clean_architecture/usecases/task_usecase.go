package usecases

import (
	"task_with_clean_arc/domain"
	"task_with_clean_arc/repository"
)

type TaskUsecase interface {
	GetTasks() ([]domain.Task, error)
	GetTaskByID(id string) (domain.Task, error)
	AddTask(task domain.Task) error
	DeleteTask(id string) error
	UpdateTask(id string, task domain.Task) error
}

type taskUsecase struct{
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase{
	return &taskUsecase{repo: repo}
}

func (u *taskUsecase) GetTasks() ([]domain.Task, error){
	return u.repo.GetAll()
}

func (u *taskUsecase) GetTaskByID(id string) (domain.Task, error){
	return u.repo.GetOne(id)
}

func (u *taskUsecase) AddTask(task domain.Task) error{
	return u.repo.Add(task)
}

func (u *taskUsecase) DeleteTask(id string) error{
	return u.repo.Delete(id)
}

func (u *taskUsecase) UpdateTask(id string, task domain.Task) error{
	return u.repo.Update(id,task)
}