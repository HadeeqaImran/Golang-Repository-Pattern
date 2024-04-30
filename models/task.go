package models

import "time"

type Status string

const (
	TODO  Status = "TODO"
	DOING Status = "DOING"
	DONE  Status = "DONE"
)

type Task struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}
