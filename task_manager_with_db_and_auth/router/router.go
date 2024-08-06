package router

import (
	"task_manager_with_db_and_auth/controllers"
	"task_manager_with_db_and_auth/middleware"

	"github.com/gin-gonic/gin"
)
func CreateRouting(){
	router := gin.Default()
	
	
	router.POST("/register",controllers.RegisterUser)
	router.POST("/login",controllers.LoginUser)

	allowed := router.Group("")
	allowed.Use(middleware.AuthUser())
	allowed.GET("/tasks", controllers.ShowAllTasks)
	allowed.GET("/tasks/:id",controllers.ShowSpecificTask)
	protected:=router.Group("/admin")
	protected.Use(middleware.AuthMiddleware("admin"))
	protected.PUT("/tasks/:id",controllers.UpdateTask)
	protected.DELETE("/tasks/:id",controllers.DeleteTask)
	protected.POST("/tasks",controllers.AddTask)
	protected.POST("/register",controllers.RegisterAdmin)
	protected.POST("/activate/:username",controllers.Activate)
	protected.POST("/deactivate/:username",controllers.DeActivate)
	protected.GET("/promote/:username",controllers.Promote)


	router.Run("localhost:8080")
}