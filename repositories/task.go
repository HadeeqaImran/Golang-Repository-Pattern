package repositories

import (
	"errors"
	"repo_pattern/domain/entities"
	"time"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

type Status string

const (
	TODO  Status = "TODO"
	DOING Status = "DOING"
	DONE  Status = "DONE"
)

type TaskRow struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	Title       string
	Description string
	Status      Status
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (TaskRow) TableName() string {
	return "tasks"
}

func NewTaskRow(task entities.CreateTask) *TaskRow {
	return &TaskRow{
		Title:       task.Title,
		Description: task.Description,
		Status:      Status(task.Status),
	}
}

func (tr *TaskRepository) Create(task *entities.CreateTask) error {
	taskRow := TaskRow{
		Title:       task.Title,
		Description: task.Description,
		Status:      Status(task.Status),
	}
	result := tr.db.Model(&TaskRow{}).Create(&taskRow)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *TaskRepository) Update(id uint64, task *entities.UpdateTask) error {
	result, err := tr.GetById(uint(id))
	if err != nil {
		return err
	}

	taskRow := TaskRow{
		ID:          result.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      Status(task.Status),
	}
	tr.db.Model(&TaskRow{}).Where("id = ?", id).Updates(taskRow)
	return nil
}

func (tr *TaskRepository) Delete(id uint) error {
	result := tr.db.Delete(&entities.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no task found with the given ID")
	}
	return nil
}

func (tr *TaskRepository) GetAll() ([]entities.Task, error) {
	var tasks []entities.Task
	result := tr.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (tr *TaskRepository) GetById(id uint) (*entities.Task, error) {
	var task entities.Task
	result := tr.db.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

// Change task status
// UpdateTaskStatus updates the status of a task.
func (tr *TaskRepository) UpdateTaskStatus(id uint, newStatus entities.StatusChangeRequest) error {
	// Get the task by ID
	task, err := tr.GetById(id)
	if err != nil {
		return err
	}

	// Update the task's status
	task.Status = newStatus.Status

	// Save the updated task
	result := tr.db.Updates(&task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
