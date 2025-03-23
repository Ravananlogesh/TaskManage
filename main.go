package main

import (
	"log"
	"os"
	"tasks/config"
	"tasks/internal/handlers/login"
	"tasks/internal/handlers/register"
	"tasks/internal/handlers/tasks"
	"tasks/internal/middleware"
	database "tasks/migrations"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	lFile, lErr := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if lErr != nil {
		log.Fatalf("error opening file:  %v", lErr)
	}
	defer lFile.Close()

	log.SetOutput(lFile)
	config.LoadGlobalConfig("toml/config.toml")
	err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Error Occur in DB Connection : ", err)
	}
	r := gin.Default()
	r.POST("/login", login.Login)
	r.POST("/register", register.Register, middleware.IPRestrictionMiddleware(), register.Register)

	taskHandler := r.Group("/tasks")
	taskHandler.Use(middleware.AuthMiddleware(), middleware.RateLimitMiddleware())
	// r.Use(middleware.RateLimitMiddleware())

	taskHandler.POST("/", tasks.CreateTask)
	taskHandler.GET("/", tasks.GetAllTask)
	taskHandler.GET("/:id", tasks.GetTaskUseByID)
	taskHandler.PUT("/:id", tasks.UpdateTask)
	taskHandler.DELETE("/:id", tasks.DeleteTask)
	log.Println("Server is running on port 1803")

	r.Run(":1803")

}
