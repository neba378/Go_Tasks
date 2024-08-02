package router

import (
	"task_manager_with_db/controllers"

	"github.com/gin-gonic/gin"
)
func CreateRouting(){
	router := gin.Default()
	router.GET("/tasks", controllers.ShowAllTasks)
	router.GET("/tasks/:id",controllers.ShowSpecificTask)
	router.PUT("/tasks/:id",controllers.UpdateTask)
	router.DELETE("tasks/:id",controllers.DeleteTask)
	router.POST("tasks/",controllers.AddTask)
	router.Run("localhost:8080")
}