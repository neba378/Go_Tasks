package handler

import (
	"net/http"
	"task_with_clean_arc/domain"
	"task_with_clean_arc/usecases"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Usecase usecases.UserUsecase
}

func NewUserHandler(u usecases.UserUsecase) *UserHandler {
	return &UserHandler{Usecase: u}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Usecase.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message":"User registered"})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erraor": err.Error()})
		return
	}
	token, err := h.Usecase.LoginUser(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erbror": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) RegisterAdmin(c *gin.Context){
	var newAdmin domain.User
	err := c.ShouldBindJSON(&newAdmin)
	if err!= nil{
		c.JSON(500, gin.H{"error":"internal error"})
		return
	}
	err = h.Usecase.RegisterAdmin(newAdmin)
	if err!=nil{
		c.JSON(http.StatusConflict, gin.H{"error":err.Error()})
		return
	}

	c.JSON(200, gin.H{"message":"Successfully registered!"})
}
func (h *UserHandler) Promote(c *gin.Context){
	username := c.Param("username")
	var UpdatedUser domain.User
	if err := c.ShouldBind(&UpdatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println(tasks)
	err := h.Usecase.UpdateUser(username)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"}) // indicates the task with given id is not found in the db
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User updated"}) // updates successfully 
}
func (h *UserHandler) Activate(c *gin.Context) {
	username := c.Param("username")
	err := h.Usecase.Activate(username)
	if err!=nil{
		c.JSON(500, gin.H{"error":err})
		return
	}
	c.JSON(200, gin.H{"message":"successfully activated!"})
}
func (h *UserHandler) DeActivate(c *gin.Context){
	username := c.Param("username")
	err := h.Usecase.DeActivate(username)
	if err!=nil{
		c.JSON(500, gin.H{"error":err})
		return
	}
	c.JSON(200, gin.H{"message":"successfully deactivated!"})

}