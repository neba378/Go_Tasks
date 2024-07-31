package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var tasks = []Task{
    {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

var LastID = 3

func ShowAllTasks (c *gin.Context){
	c.IndentedJSON(http.StatusOK,tasks)
}

func ShowSpecificTask(c *gin.Context){
	id := c.Param("id")
	for _,t := range tasks{
		if t.ID == id{
			c.IndentedJSON(http.StatusOK,t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"No such task found!"})
}

func UpdateTask (c *gin.Context){
	id := c.Param("id")
	var UpdatedTask Task
	if err := c.ShouldBind(&UpdatedTask); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, task := range tasks{
		if task.ID == id{
			if UpdatedTask.Title != ""{
				tasks[i].Title = UpdatedTask.Title

			}
			if UpdatedTask.Description != ""{
				tasks[i].Description = UpdatedTask.Description
			}
			
			c.IndentedJSON(http.StatusOK, gin.H{"updated Task":tasks[i]})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not found"})
}

func DeleteTask(c *gin.Context) {
    id := c.Param("id")

    for i, val := range tasks {
        if val.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Task removed"})
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func AddTask(c *gin.Context) {
    var newTask Task

    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	newTask.Status = "Pending"
	LastID+=1
	t := strconv.Itoa(LastID)
	newTask.ID = t
	newTask.DueDate = time.Now()
    tasks = append(tasks, newTask)
    c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func main(){
	fmt.Print("Task Manager API")
	router := gin.Default()
	router.GET("/tasks", ShowAllTasks)
	router.GET("/tasks/:id",ShowSpecificTask)
	router.PUT("/tasks/:id",UpdateTask)
	router.DELETE("tasks/:id",DeleteTask)
	router.POST("tasks/",AddTask)
	router.Run("localhost:8080")

}