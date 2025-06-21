package entity

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusProcessing TaskStatus = "processing"
	StatusCompleted  TaskStatus = "completed"
	StatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID          string
	Status      TaskStatus
	CreatedAt   time.Time
	StartedAt   *time.Time
	CompletedAt *time.Time
	Result      *string
	Error       *string
	DurationMs  int64
}

func (t *Task) ProcessingTime() time.Duration {
	if t.StartedAt == nil {
		return 0
	}

	endTime := time.Now()
	if t.CompletedAt != nil {
		endTime = *t.CompletedAt
	}

	return endTime.Sub(*t.StartedAt)
}

func GenerateID() string {
	return uuid.NewString()
}
