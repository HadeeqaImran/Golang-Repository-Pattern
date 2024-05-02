package entities

import "time"

type Status string

const (
	TODO  Status = "TODO"
	DOING Status = "DOING"
	DONE  Status = "DONE"
)

type Task struct {
	ID          uint
	CreatedAt   time.Time
	Title       string
	Description string
	Status      Status
}

type CreateTask struct {
	Title       string
	Description string
	Status      Status
}

type UpdateTask struct {
	Title       string
	Description string
	Status      Status
}

type StatusChangeRequest struct {
	Status Status
}
