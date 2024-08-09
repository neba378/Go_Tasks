package router

import (
	"task_with_clean_arc/Delivery/controllers"
	"task_with_clean_arc/infrastructures"
	"task_with_clean_arc/repository"
	"task_with_clean_arc/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)
func CreateRouting(client *mongo.Client) {
    router := gin.Default()

    // Initialize repositories
    userRepo := repository.NewUserRepository(client)
    taskRepo := repository.NewTaskRepository(client)

    // Initialize use cases
    userUsecase := usecases.NewUserUsecase(userRepo)
    taskUsecase := usecases.NewTaskUsecase(taskRepo)

    // Initialize handlers
    userHandler := controllers.NewUserHandler(userUsecase)
    taskHandler := controllers.NewTaskHandler(taskUsecase)

    // Public routes
    router.POST("/register", userHandler.RegisterUser)
    router.POST("/login", userHandler.LoginUser)

    // Routes for authenticated users
    allowed := router.Group("")
    allowed.Use(infrastructures.AuthUser())
    allowed.GET("/tasks", taskHandler.GetTasks)
    allowed.GET("/tasks/:id", taskHandler.GetTaskByID)

    // Routes for admin users
    protected := router.Group("/admin")
    protected.Use(infrastructures.AuthMiddleware("admin"))
    protected.PUT("/tasks/:id", taskHandler.UpdateTask)
    protected.DELETE("/tasks/:id", taskHandler.DeleteTask)
    protected.POST("/tasks", taskHandler.AddTask)
    protected.POST("/register", userHandler.RegisterAdmin)
    protected.POST("/activate/:username", userHandler.Activate)
    protected.POST("/deactivate/:username", userHandler.DeActivate)
    protected.GET("/promote/:username", userHandler.Promote)

    // Run the server
    router.Run("localhost:8080")
}