package entity

import "time"

type Task struct {
	ID          int
	Name        string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
