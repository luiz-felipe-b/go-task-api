package tasks

import (
	"to-do-api/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	dbConn := db.Connect()
	repo := &TaskRepository{DB: dbConn}
	service := &TaskService{Repo: repo}
	handler := &TaskHandler{Service: service}

	routes := router.Group("/tasks")
	{
		routes.GET("/", handler.GetAllTasks)
		routes.POST("/", handler.CreateTask)
	}
}

type TaskHandler struct {
	Service *TaskService
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.Service.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
