package database

import (
	"fmt"
	"log"
	"tasks/config"
	"tasks/internal/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GDB *gorm.DB

func ConnectDatabase() error {
	cf := config.GetConfig()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cf.Database.Host, cf.Database.Port, cf.Database.User,
		cf.Database.Pass, cf.Database.Name, cf.Database.Sslmode)
	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	err = TaskStatusEnum(db)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	err = db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	GDB = db
	log.Println("Database connection established successfully")
	return nil
}

func GetDB() *gorm.DB {
	return GDB
}
func TaskStatusEnum(db *gorm.DB) error {
	var exists bool
	checkEnumQuery := `
		SELECT EXISTS (
			SELECT 1 FROM pg_type WHERE typname = 'task_status'
		);
	`
	if err := db.Raw(checkEnumQuery).Scan(&exists).Error; err != nil {
		return err
	}

	if !exists {
		createEnumQuery := `CREATE TYPE task_status AS ENUM ('Pending', 'In Progress', 'Completed');`
		return db.Exec(createEnumQuery).Error
	}
	return nil
}
