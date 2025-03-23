package request

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required,oneof=Pending 'In Progress' Completed"`
	DueDate     string `json:"due_date,omitempty"`
}

func (req *CreateTaskRequest) Validate() error {
	return Validate.Struct(req)
}

type UpdateTaskRequest struct {
	Title       string     `json:"title,omitempty" validate:"omitempty,min=3,max=100"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status,omitempty" validate:"omitempty,oneof=Pending 'In Progress' Completed"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

func (req *UpdateTaskRequest) Validate() error {
	return validateAndFormat(req)
}
func validateAndFormat(req any) error {
	err := Validate.Struct(req)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	errorMessages := make([]string, len(validationErrors))

	for i, fieldErr := range validationErrors {
		jsonField := fieldErr.Field()
		if jsonTag := fieldErr.StructField(); jsonTag != "" {
			jsonField = fieldErr.StructNamespace()
		}
		errorMessages[i] = fmt.Sprintf("%s: %s", jsonField, fieldErr.Error())
	}

	return fmt.Errorf("validation failed: %s", errorMessages)
}
