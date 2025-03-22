package repo

import (
	"tasks/internal/models"
	database "tasks/migrations"

	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByName(username string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

func NewUserRepo() UserRepository {
	return &userRepo{DB: database.GetDB()}
}

func (r *userRepo) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepo) GetUserByName(username string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
