package controllers

import (
	"net/http"
	"task_manager_with_db/data"
	"task_manager_with_db/models"

	"github.com/gin-gonic/gin"
)

var tasks []models.Task = []models.Task{}



func ShowAllTasks(c *gin.Context) {
	tasks = data.GetAllTasks()
	// fmt.Print(tasks)
	c.IndentedJSON(http.StatusOK, tasks) // gives JSON formatted list of tasks to the end point
}

func ShowSpecificTask(c *gin.Context) {
	id := c.Param("id")
	tasks = data.GetAllTasks()
	res, err := data.GetOne(id)

	if err!= nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err}) 
	}
	c.IndentedJSON(http.StatusOK, res) // gives JSON formatted task to the end point (the specific task needed)
	
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var UpdatedTask models.Task
	if err := c.ShouldBind(&UpdatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println(tasks)
	err := data.UpdateTask(id,UpdatedTask)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"}) // indicates the task with given id is not found in the db
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated"}) // updates successfully 
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := data.DeleteTask(id) // the function is from data/task_service.go
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"message": "Task is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}

func AddTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := data.AddTask(newTask)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created"}) // a green light for task creation
}