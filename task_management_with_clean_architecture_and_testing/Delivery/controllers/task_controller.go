package controllers

import (
	"fmt"
	"net/http"
	"task_with_clean_arc_and_test/domain"
	"task_with_clean_arc_and_test/usecases"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	usecase usecases.TaskUsecase
}

func NewTaskHandler(usecase usecases.TaskUsecase) *TaskHandler {
	return &TaskHandler{usecase: usecase}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.usecase.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks) // Ensure that only the tasks slice is returned
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.usecase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) AddTask(c *gin.Context) {
	var newTask domain.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.AddTask(newTask)
	if err != nil {
		fmt.Print("ufff", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := h.usecase.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task not found"})
		return // Return early to avoid sending a success response
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted!"})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.usecase.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully updated!"})
}
