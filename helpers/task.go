package helpers

import (
	"Checkmarx/models"
	"fmt"
)

func ValidateTaskFields(field models.Task) (bool, error) {
	if field.Title == "" {
		return false, fmt.Errorf("title is required")
	}
	if field.Description == "" {
		return false, fmt.Errorf("description is required")
	}
	if field.Status == "" {
		return false, fmt.Errorf("status is required")
	}
	return true, nil
}

func ValidateTaskUpdate(updatedTask models.Task) (bool, error) {
	if updatedTask.Title == "" && updatedTask.Description == "" && updatedTask.Status == "" {
		return false, fmt.Errorf("at least one field (Title, Description, Status) must be updated")
	}
	return true, nil
}
