package model

import (
	"time"
)

type Task struct {
	ID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Title string
}