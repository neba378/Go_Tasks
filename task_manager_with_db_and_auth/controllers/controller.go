package controllers

import (
	"net/http"
	"task_manager_with_db_and_auth/data"
	"task_manager_with_db_and_auth/models"

	"github.com/dgrijalva/jwt-go"
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


// User functions

func RegisterUser(c *gin.Context){
	var newUser models.User
	err := c.ShouldBindJSON(&newUser)
	if err!= nil{
		c.JSON(500, gin.H{"error":"internal error"})
		return
	}
	// role:=c.Param("role")
	// fmt.Print("-=-=-=",role)
	err = data.Register(newUser)
	if err!=nil{
		c.JSON(http.StatusConflict, gin.H{"error":err.Error()})
		return
	}

	c.JSON(200, gin.H{"message":"Successfully registered!"})
}

func LoginUser(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user); err!=nil{
		c.JSON(400, gin.H{"error":"incorrect input"})
		return
	}
	
	token, err := data.CheckUser(user)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":"Successfully logged in!",
		"token":token,
	})
}

// this function is used to handle the protected path
func ProtectedHandler(c *gin.Context){
	claims, _ := c.Get("user") 
	userClaim:= claims.(jwt.MapClaims)
	userID := userClaim["id"]
	username:=userClaim["username"]

	response := gin.H{
		"message":"You have successfully accessed the protected route!",
		"id":userID,
		"username":username,
	}

	c.JSON(200, response)
}


func RegisterAdmin(c *gin.Context){
	var newAdmin models.User
	err := c.ShouldBindJSON(&newAdmin)
	if err!= nil{
		c.JSON(500, gin.H{"error":"internal error"})
		return
	}
	err = data.RegisterAdmin(newAdmin)
	if err!=nil{
		c.JSON(http.StatusConflict, gin.H{"error":err.Error()})
		return
	}

	c.JSON(200, gin.H{"message":"Successfully registered!"})
}

func Activate(c *gin.Context){
	username := c.Param("username")
	err := data.Activate(username)
	if err!=nil{
		c.JSON(500, gin.H{"error":err})
		return
	}
	c.JSON(200, gin.H{"message":"successfully activated!"})
}

func DeActivate(c *gin.Context){
	username := c.Param("username")
	err := data.DeActivate(username)
	if err!=nil{
		c.JSON(500, gin.H{"error":err})
		return
	}
	c.JSON(200, gin.H{"message":"successfully deactivated!"})
}