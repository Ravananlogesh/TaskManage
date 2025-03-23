package repo

import (
	"tasks/internal/models"
	"tasks/internal/utils"
	database "tasks/migrations"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id uint) error
	GetTasksWithFilter(log *utils.Logger, filter *models.TaskFilter) (*[]models.Task, error)
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
func (r *taskRepo) GetTasksWithFilter(log *utils.Logger, filter *models.TaskFilter) (*[]models.Task, error) {
	log.Log(utils.INFO, "GetTasksWithFilter +")
	var taskArr []models.Task
	query := r.DB.Model(&models.Task{})
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.DueDateAfter != "" {
		query = query.Where("due_date >= ?", filter.DueDateAfter)
	}
	if filter.DueDateBefore != "" {
		query = query.Where("due_date <= ?", filter.DueDateBefore)
	}
	query = query.Order(filter.SortBy + " " + filter.SortOrder)
	query = query.Offset((filter.Page - 1) * filter.Limit).Limit(filter.Limit)

	err := query.Find(&taskArr).Error
	if err != nil {
		return nil, err
	}
	log.Log(utils.INFO, "GetTasksWithFilter -")

	return &taskArr, nil

}
