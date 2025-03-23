package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"tasks/internal/handlers/login"
	"tasks/internal/handlers/tasks"
	"tasks/internal/models"
	"tasks/internal/models/request"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	r := gin.Default()
	r.POST("/tasks", tasks.CreateTask)

	reqBody := request.CreateTaskRequest{
		Title:       "test task",
		Description: "description of test task",
		Status:      "pending",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetAllTask(t *testing.T) {
	r := gin.Default()
	r.GET("/tasks", tasks.GetAllTask)

	req, _ := http.NewRequest("GET", "/tasks?page=1&limit=10", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetTaskUseByID(t *testing.T) {
	r := gin.Default()
	r.GET("/tasks/:id", tasks.GetTaskUseByID)

	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestLogin(t *testing.T) {
	r := gin.Default()
	r.POST("/login", login.Login)

	loginReq := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	body, _ := json.Marshal(loginReq)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
func TestUpdateTask(t *testing.T) {
	r := gin.Default()
	r.PUT("/tasks/:id", tasks.UpdateTask)

	reqBody := models.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "completed",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteTask(t *testing.T) {
	r := gin.Default()
	r.DELETE("/tasks/:id", tasks.DeleteTask)

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
