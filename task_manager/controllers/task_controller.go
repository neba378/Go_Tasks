package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

var tasks = data.GetAllTasks()

func ShowAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks) // gives JSON formatted list of tasks to the end point
}

func ShowSpecificTask(c *gin.Context) {
	id := c.Param("id")
	for _, t := range tasks {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	// gives JSON formatted task to the end point (the specific task needed)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such task found!"}) 
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var UpdatedTask models.Task
	if err := c.ShouldBind(&UpdatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println(tasks)
	NewTasks,err := data.UpdateTask(id,UpdatedTask,tasks)
	if err!=nil{
		// fmt.Println(err)
		// unsuccessful try to update
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"}) 
		return
	}
	tasks = NewTasks // re assign the tasks after the update
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	NewTasks,err := data.DeleteTask(id,tasks) // the function is from data/task_service.go
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	tasks = NewTasks
	c.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}

func AddTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	tasks = data.AddTask(newTask,tasks)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}