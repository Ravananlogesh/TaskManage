package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JSONResponse(c *gin.Context, status int, success bool, message string, data any) {
	c.JSON(status, Response{
		Success: success,
		Message: message,
		Data:    data,
	})
}
