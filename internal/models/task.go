package models

import (
	"time"

	"gorm.io/gorm"
)

// TaskStatus defines possible task statuses
type TaskStatus string

const (
	Pending    TaskStatus = "Pending"
	InProgress TaskStatus = "In Progress"
	Completed  TaskStatus = "Completed"
)

// Task represents a task entity
type Task struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `gorm:"type:varchar(20);default:'Pending'" json:"status"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	UserID      uint       `gorm:"not null;constraint:OnDelete:CASCADE;" json:"user_id"`
}

func (t *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return
}
