package service

import (
	"Checkmarx/models"
	"Checkmarx/repository"
)

type TaskService struct {
	repo *repository.TaskManager
}

func NewTaskService(repo *repository.TaskManager) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) AddTask(task models.Task) (models.Task, error) {
	return s.repo.AddTask(task)
}

func (s *TaskService) GetTask(id int) (models.Task, error) {
	return s.repo.GetTask(id)
}

func (s *TaskService) GetAllTasks() map[int]models.Task {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTask(id int, updatedTask models.Task) (models.Task, error) {
	return s.repo.Update(id, updatedTask)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
