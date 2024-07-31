package data

import (
	"errors"
	"strconv"
	"task_manager/models"
	"time"
)

var tasks = []models.Task{}
var LastID = 1

// this function is used to provide the whole task
func GetAllTasks() []models.Task {
    return tasks
}
func AddTask(newTask models.Task, tasks []models.Task) []models.Task{
	newTask.Status = "Pending"
	LastID+=1
	t := strconv.Itoa(LastID)
	newTask.ID = t
	newTask.DueDate = time.Now()
    tasks = append(tasks, newTask)
	return tasks // returns tasks after successful addition of new task
	
}



func DeleteTask(id string, tasks []models.Task) ([]models.Task,error) {
	for i, val := range tasks{
        if val.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            return tasks, nil // returns the task after successful deletion of the task
        }
    }
	return tasks, errors.New("not found") // indicates no such task exist
}


func UpdateTask(id string,UpdatedTask models.Task,tasks []models.Task)([]models.Task,error){
	// fmt.Println(tasks)
	for i, task := range tasks{
		if task.ID == id{
			if UpdatedTask.Title != ""{
				tasks[i].Title = UpdatedTask.Title

			}
			if UpdatedTask.Description != ""{
				tasks[i].Description = UpdatedTask.Description
			}
			return tasks,nil // returns the updated tasks if the task is in there 
		}
	}
	return tasks,errors.New("task not found!") // error indicating the task is not found
}