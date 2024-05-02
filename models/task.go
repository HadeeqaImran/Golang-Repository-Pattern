package models

import "time"

type Status string

const (
	TODO  Status = "TODO"
	DOING Status = "DOING"
	DONE  Status = "DONE"
)

type Task struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	Title       string
	Description string
	Status      Status
}
