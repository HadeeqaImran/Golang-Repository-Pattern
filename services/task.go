package services

import (
	"errors"

	"repo_pattern/dto"
	"repo_pattern/models"
	"repo_pattern/repositories"
)

type TaskService struct {
	taskRepository *repositories.TaskRepository
}

func NewTaskService(taskRepo *repositories.TaskRepository) *TaskService {
	return &TaskService{taskRepository: taskRepo}
}

func (ts *TaskService) CreateTask(task *models.Task) error {
	if !isValidStatus(task.Status) {
		return errors.New("invalid status value")
	}
	return ts.taskRepository.Create(task)
}

func (ts *TaskService) UpdateTask(task *models.Task) error {
	if !isValidStatus(task.Status) {
		return errors.New("invalid status value")
	}
	return ts.taskRepository.Update(task)
}

func (ts *TaskService) DeleteTask(id uint) error {
	// Perform any business logic validations here
	return ts.taskRepository.Delete(id)
}

func (ts *TaskService) GetAllTasks() ([]models.Task, error) {
	return ts.taskRepository.GetAll()
}

func (ts *TaskService) GetTaskById(id uint) (*models.Task, error) {
	return ts.taskRepository.GetById(id)
}

// ChangeTaskStatus changes the status of a task.
func (ts *TaskService) ChangeTaskStatus(id uint, newStatus dto.StatusRequest) error {
	// Check if the task exists
	_, err := ts.taskRepository.GetById(id)
	if err != nil {
		return err
	}

	if !isValidStatus(models.Status(newStatus.Status)) {
		return errors.New("invalid status value")
	}

	// Update the task status
	err = ts.taskRepository.UpdateTaskStatus(id, models.Status(newStatus.Status))
	if err != nil {
		return err
	}

	return nil
}

// Utility Functions
func isValidStatus(status models.Status) bool {
	switch status {
	case models.TODO, models.DOING, models.DONE:
		return true
	default:
		return false
	}
}
