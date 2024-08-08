package handler

import (
	"net/http"
	"task_with_clean_arc/domain"
	"task_with_clean_arc/usecases"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	usecase usecases.TaskUsecase
}

func NewTaskHandler(usecase usecases.TaskUsecase) *TaskHandler{
	return &TaskHandler{usecase: usecase}
}

func (h *TaskHandler) GetTasks(c *gin.Context){
	tasks,err := h.usecase.GetTasks()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context){
	id := c.Param("id")
	task,err := h.usecase.GetTaskByID(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
    c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) AddTask(c *gin.Context){
	var newTask domain.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := h.usecase.AddTask(newTask)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created"}) 
}

func (h *TaskHandler) DeleteTask(c *gin.Context){
	id := c.Param("id")
	err := h.usecase.DeleteTask(id)
	if err!=nil{
		c.JSON(500, gin.H{"err":err})
	}
	c.JSON(200, gin.H{"message":"successfully deleted!"})
}

func (h *TaskHandler) UpdateTask(c *gin.Context){
	id := c.Param("id")
	var UpdatedTask domain.Task
	if err := c.ShouldBind(&UpdatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.usecase.UpdateTask(id,UpdatedTask)
	if err!=nil{
		c.JSON(500, gin.H{"err":err})
	}
	c.JSON(200, gin.H{"message":"successfully updated!"})
}
