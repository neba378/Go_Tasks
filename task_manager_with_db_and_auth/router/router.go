package router

import (
	"task_manager_with_db_and_auth/controllers"
	"task_manager_with_db_and_auth/middleware"

	"github.com/gin-gonic/gin"
)
func CreateRouting(){
	router := gin.Default()
	router.GET("/tasks", controllers.ShowAllTasks)
	router.GET("/tasks/:id",controllers.ShowSpecificTask)
	router.PUT("/tasks/:id",controllers.UpdateTask)
	router.DELETE("/tasks/:id",controllers.DeleteTask)
	router.POST("/tasks",controllers.AddTask)
	router.POST("/register/:role",controllers.RegisterUser)
	router.POST("/login",controllers.LoginUser)


	protected:=router.Group("/protected")
	protected.Use(middleware.AuthMiddleware("admin"))
	protected.GET("/secrete",controllers.ProtectedHandler)


	router.Run("localhost:8080")
}