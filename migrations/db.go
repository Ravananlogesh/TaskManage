package database

import (
	"log"
	"os"
	"tasks/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GDB *gorm.DB

func ConnectDatabase() error {
	// var cf models.Config

	// err := config.LoadTOML("config.toml", &cf)
	// if err != nil {
	// 	return err
	// }
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cf.Database.Host, cf.Database.Port, cf.Database.User,
	// 	cf.Database.Pass, cf.Database.Name, cf.Database.Sslmode)

	// // Connect to PostgreSQL
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// 	return err
	// }
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	err = db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
		return err
	}

	GDB = db
	log.Println("Database connection established successfully")
	return nil
}

func GetDB() *gorm.DB {
	return GDB
}
