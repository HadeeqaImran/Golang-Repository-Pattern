package services

import (
	"errors"

	"repo_pattern/domain/entities"
	"repo_pattern/repositories"
)

type Status string

const (
	TODO  Status = "TODO"
	DOING Status = "DOING"
	DONE  Status = "DONE"
)

type TaskServiceStruct struct {
	Title       string
	Description string
	Status      Status
}

type StatusChangeStruct struct {
	Status Status
}

type TaskService struct {
	taskRepository *repositories.TaskRepository
}

func NewTaskService(taskRepo *repositories.TaskRepository) *TaskService {
	return &TaskService{taskRepository: taskRepo}
}

func (ts *TaskServiceStruct) NewTaskEntity() *entities.CreateTask {
	return &entities.CreateTask{
		Title:       ts.Title,
		Description: ts.Description,
		Status:      entities.Status(ts.Status),
	}
}

func (ts *TaskServiceStruct) NewUpdateTaskEntity() *entities.UpdateTask {
	return &entities.UpdateTask{
		Title:       ts.Title,
		Description: ts.Description,
		Status:      entities.Status(ts.Status),
	}
}

func (sc *StatusChangeStruct) NewStatusChangeEntity() *entities.StatusChangeRequest {
	return &entities.StatusChangeRequest{
		Status: entities.Status(sc.Status),
	}
}

func (ts *TaskService) CreateTask(task *TaskServiceStruct) error {
	if !isValidStatus(string(task.Status)) {
		return errors.New("invalid status value")
	}
	return ts.taskRepository.Create(task.NewTaskEntity())
}

func (ts *TaskService) UpdateTask(id uint64, task *TaskServiceStruct) error {
	if !isValidStatus(string(task.Status)) {
		return errors.New("invalid status value")
	}
	return ts.taskRepository.Update(id, task.NewUpdateTaskEntity())
}

func (ts *TaskService) DeleteTask(id uint) error {
	return ts.taskRepository.Delete(id)
}

func (ts *TaskService) GetAllTasks() ([]entities.Task, error) {
	return ts.taskRepository.GetAll()
}

func (ts *TaskService) GetTaskById(id uint) (*entities.Task, error) {
	return ts.taskRepository.GetById(id)
}

// ChangeTaskStatus changes the status of a task.
func (ts *TaskService) ChangeTaskStatus(id uint, newStatus *StatusChangeStruct) error {
	// Check if the task exists
	_, err := ts.taskRepository.GetById(id)
	if err != nil {
		return err
	}

	if !isValidStatus(string(newStatus.Status)) {
		return errors.New("invalid status value")
	}

	// Update the task status
	err = ts.taskRepository.UpdateTaskStatus(id, *newStatus.NewStatusChangeEntity())
	if err != nil {
		return err
	}

	return nil
}

// Utility Functions
func isValidStatus(status string) bool {
	switch status {
	case "TODO", "DOING", "DONE":
		return true
	default:
		return false
	}
}
