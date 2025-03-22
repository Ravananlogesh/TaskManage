package main

import (
	"log"
	"os"
	"tasks/internal/middleware"
	database "tasks/migrations"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	lFile, lErr := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if lErr != nil {
		log.Fatalf("error opening file:  %v", lErr)
	}
	defer lFile.Close()

	log.SetOutput(lFile)

	err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Error Occur in DB Connection : ", err)
	}
	r := gin.Default()
	r.POST("/login", login.Login)
	r.Use(middleware.AuthMiddleware(), middleware.RateLimitMiddleware())
	taskHandler := r.Group("/tasks")

	taskHandler.POST()

}
