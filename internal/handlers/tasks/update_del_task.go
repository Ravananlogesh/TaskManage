package tasks

import (
	"net/http"
	"strconv"
	"tasks/internal/models"
	"tasks/internal/repo"
	"tasks/internal/utils"

	"github.com/gin-gonic/gin"
)

func UpdateTask(c *gin.Context) {
	log := new(utils.Logger)
	log.SetSid(c.Request)
	taskId, _ := strconv.Atoi(c.Param("id"))
	var req models.Task
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Log(utils.ERROR, "UT001", err.Error())
		utils.JSONResponse(c, http.StatusBadRequest, false, "invalid request "+err.Error(), nil)
		return
	}
	taskRepo := repo.NewTaskRepo()
	task, err := taskRepo.GetTaskByID(uint(taskId))
	if err != nil {
		log.Log(utils.ERROR, "UT001", err.Error())
		utils.JSONResponse(c, http.StatusBadRequest, false, "invalid request "+err.Error(), nil)
		return
	}
	task.Title = req.Title
	task.Description = req.Description
	if req.DueDate != "" {
		task.DueDate = req.DueDate
	}
	task.Status = req.Status

	err = taskRepo.UpdateTask(task)
	if err != nil {
		log.Log(utils.ERROR, "UT001", err.Error())
		utils.JSONResponse(c, http.StatusBadRequest, false, "invalid request "+err.Error(), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, true, "Task updated successfully", task)
}
func DeleteTask(c *gin.Context) {
	log := new(utils.Logger)
	log.SetSid(c.Request)
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Log(utils.ERROR, "DT001", err.Error())
		utils.JSONResponse(c, http.StatusBadRequest, false, "invalid request "+err.Error(), nil)
		return
	}
	taskRepo := repo.NewTaskRepo()
	err = taskRepo.DeleteTask(uint(taskID))
	if err != nil {
		log.Log(utils.ERROR, "DT002", err.Error())
		utils.JSONResponse(c, http.StatusNotFound, false, "Task not found", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, true, "Task successfully deleted", nil)
}
