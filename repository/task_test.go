package repository

import (
	"Checkmarx/helpers"
	"Checkmarx/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTaskRepository(t *testing.T) {
	repo := NewTaskManager()

	newTaskPositive := models.Task{
		Title:       "Test Task",
		Description: "This is a test task.",
		Status:      "Pending",
		CreatedAt:   time.Now(),
	}

	_, err := repo.AddTask(newTaskPositive)
	assert.NoError(t, err, "unexpected error while adding task")
	assert.Equal(t, 1, repo.counter, "expected ID 1")

	task, err := repo.GetTask(repo.counter)

	assert.NoError(t, err, "unexpected error while getting task")
	assert.Equal(t, "Test Task", task.Title, "expected title 'Test Task'")

	updatedTaskData := models.Task{
		Title:       "Updated Task",
		Description: "This is a test task.",
		Status:      "Completed",
	}
	updatedTask, err := repo.Update(repo.counter, updatedTaskData)

	assert.NoError(t, err, "unexpected error while updating task")
	assert.Equal(t, "Updated Task", updatedTask.Title, "expected title 'Updated Task'")
	assert.Equal(t, "Completed", updatedTask.Status, "expected status 'Completed'")
	assert.Equal(t, "This is a test task.", updatedTask.Description, "expected description to remain the same")

	allTasks := repo.GetAllTasks()
	assert.Equal(t, 1, len(allTasks), "expected 1 task")

	err = repo.Delete(repo.counter)
	assert.NoError(t, err, "unexpected error while deleting task")

	_, err = repo.GetTask(repo.counter)
	assert.Error(t, err, "expected error, got none")

}

func TestValidateAddTask(t *testing.T) {
	tests := []struct {
		name     string
		task     models.Task
		expected bool
		errMsg   string
	}{
		{
			name: "Valid task",
			task: models.Task{
				Title:       "Test Task",
				Description: "Test Description",
				Status:      "open",
				CreatedAt:   time.Now(),
			},
			expected: true,
			errMsg:   "",
		},
		{
			name: "Empty title",
			task: models.Task{
				Title:       "",
				Description: "Test Description",
				Status:      "open",
				CreatedAt:   time.Now(),
			},
			expected: false,
			errMsg:   "title is required",
		},
		{
			name: "Empty description",
			task: models.Task{
				Title:       "Test Task",
				Description: "",
				Status:      "open",
				CreatedAt:   time.Now(),
			},
			expected: false,
			errMsg:   "description is required",
		},
		{
			name: "Empty status",
			task: models.Task{
				Title:       "Test Task",
				Description: "Test Description",
				Status:      "",
				CreatedAt:   time.Now(),
			},
			expected: false,
			errMsg:   "status is required",
		},
		{
			name: "Empty title, description and status",
			task: models.Task{
				Title:       "",
				Description: "",
				Status:      "",
				CreatedAt:   time.Now(),
			},
			expected: false,
			errMsg:   "title is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := helpers.ValidateTaskFields(tt.task)
			assert.Equal(t, tt.expected, valid, "expected valid: %v, got %v", tt.expected, valid)
			if !(tt.expected) {
				assert.EqualError(t, err, tt.errMsg, "expected error message: %v, got: %v", tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err, "unexpected error for valid task")
			}
		})
	}
}

func TestValidateUpdateTask(t *testing.T) {
	tests := []struct {
		name     string
		task     models.Task
		expected bool
		errMsg   string
	}{
		{
			name: "Valid task",
			task: models.Task{
				Title:       "Test Task",
				Description: "Test Description",
				Status:      "open",
				CreatedAt:   time.Now(),
			},
			expected: true,
			errMsg:   "",
		},
		{
			name: "Empty title, description and status",
			task: models.Task{
				Title:       "",
				Description: "",
				Status:      "",
				CreatedAt:   time.Now(),
			},
			expected: false,
			errMsg:   "at least one field (Title, Description, Status) must be updated",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := helpers.ValidateTaskUpdate(tt.task)
			assert.Equal(t, tt.expected, valid, "expected valid: %v, got %v", tt.expected, valid)
			if !(tt.expected) {
				assert.EqualError(t, err, tt.errMsg, "expected error message: %v, got: %v", tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err, "expected no error for valid task")
			}
		})
	}
}
