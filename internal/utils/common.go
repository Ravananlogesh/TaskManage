package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JSONResponse(c *gin.Context, status int, success bool, message string, data any) {
	logEntry := Log.WithFields(logrus.Fields{
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
		"ip":     c.ClientIP(),
		"status": status,
	})

	if status >= 400 {
		logEntry.Error(message)
	} else {
		logEntry.Info(message)
	}

	c.JSON(status, Response{
		Success: success,
		Message: message,
		Data:    data,
	})
}

func JSONErrorResponse(c *gin.Context, statusCode int, err error) {
	if err == nil {
		Log.Warn("JSONErrorResponse called with nil error")
		err = gin.Error{Err: errors.New("Unknown error")}
	}

	Log.WithFields(logrus.Fields{
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
		"ip":     c.ClientIP(),
		"status": statusCode,
		"error":  err.Error(),
	}).Error("Request failed")

	c.JSON(statusCode, gin.H{"error": err.Error()})
}
