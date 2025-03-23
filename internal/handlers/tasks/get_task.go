package tasks

import (
	"net/http"
	"strconv"
	"tasks/internal/models"
	"tasks/internal/repo"
	"tasks/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetAllTask(c *gin.Context) {
	log := new(utils.Logger)
	log.SetSid(c.Request)
	log.Log("GetAllTask started...")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Log(utils.ERROR, "GAT001", err.Error())
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Failed to fetch tasks", nil)
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		log.Log(utils.ERROR, "GAT002", err.Error())
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Failed to fetch tasks", nil)
		return
	}
	status := c.Query("status")
	dueDateAfter := c.Query("due_date_after")
	dueDateBefore := c.Query("due_date_before")
	sortBy := c.DefaultQuery("sort_by", "created_at")
	sortOrder := c.DefaultQuery("sort_order", "desc")

	filter := models.TaskFilter{
		Page:          page,
		Limit:         limit,
		Status:        status,
		DueDateAfter:  dueDateAfter,
		DueDateBefore: dueDateBefore,
		SortBy:        sortBy,
		SortOrder:     sortOrder,
	}
	taskRepo := repo.NewTaskRepo()
	log.Log(utils.INFO, "TaskFilter Details", filter)
	tasks, err := taskRepo.GetTasksWithFilter(log, &filter)
	if err != nil {
		log.Log(utils.ERROR, "GAT003", err.Error())
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Failed to fetch tasks", nil)
	}
	log.Log("GetAllTask end...")

	utils.JSONResponse(c, http.StatusOK, true, "task fetch sucessfully", &tasks)

}

func GetTaskUseByID(c *gin.Context) {
	log := new(utils.Logger)
	log.SetSid(c.Request)
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Log(utils.ERROR, "GT001", err.Error())
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Failed to fetch tasks", nil)
		return
	}
	taskRepo := repo.NewTaskRepo()

	task, err := taskRepo.GetTaskByID(uint(taskId))
	if err != nil {
		log.Log(utils.ERROR, "GT002", err.Error())
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Failed to fetch tasks", nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, true, "Task fetched successfully", task)

}
