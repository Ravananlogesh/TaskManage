package repo

import (
	"tasks/internal/models"
	database "tasks/migrations"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id uint) error
}

type taskRepo struct {
	DB *gorm.DB
}

func NewTaskRepo() TaskRepository {
	return &taskRepo{DB: database.GetDB()}
}

func (r *taskRepo) CreateTask(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *taskRepo) GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := r.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func (r *taskRepo) UpdateTask(task *models.Task) error {
	return r.DB.Save(task).Error
}

func (r *taskRepo) DeleteTask(id uint) error {
	return r.DB.Delete(&models.Task{}, id).Error
}
