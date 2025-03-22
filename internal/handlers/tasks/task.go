package tasks

import (
	"net/http"
	"tasks/internal/models"
	"tasks/internal/models/request"
	"tasks/internal/repo"
	"tasks/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	log := new(utils.Logger)
	log.SetSid(c.Request)
	var task models.Task
	var req request.CreateTaskRequest

	log.Log("CreateTask started")
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Log(utils.ERROR, "CT001", err.Error())
		utils.JSONResponse(c, http.StatusBadRequest, false, "Invalid request "+err.Error(), nil)
	}

	if err := req.Validate(); err != nil {
		log.Log(utils.ERROR, "CT002", err.Error())
		utils.JSONResponse(c, http.StatusBadRequest, false, "Invalid request "+err.Error(), nil)
	}

	task = models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      models.TaskStatus(req.Status),
		DueDate:     req.DueDate,
	}
	taskRepo := repo.NewTaskRepo()
	err := taskRepo.CreateTask(&task)
	if err != nil {
		log.Log(utils.ERROR, "CT003", err.Error())
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Faild to create task "+err.Error(), nil)
	}

	log.Log("CreateTask end")
	utils.JSONResponse(c, http.StatusOK, true, "Task created successfully", task)
}
