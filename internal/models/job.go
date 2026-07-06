package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type JobStatus string

const (
	StatusQueued  JobStatus = "queued"
	StatusRunning JobStatus = "running"
	StatusDone    JobStatus = "done"
	StatusFailed  JobStatus = "failed"
	StatusDead    JobStatus = "dead"
)

type Job struct {
	ID             uuid.UUID       `json:"id"`
	QueueName      string          `json:"queue_name"`
	Payload        json.RawMessage `json:"payload"`
	JobType        string          `json:"job_type"`
	Status         JobStatus       `json:"status"`
	Priority       int             `json:"priority"`
	Attempts       int             `json:"attempts"`
	MaxAttempts    int             `json:"max_attempts"`
	RunAt          time.Time       `json:"run_at"`
	LockedBy       *string         `json:"locked_by,omitempty"`
	LockedUntil    *time.Time      `json:"locked_until,omitempty"`
	IdempotencyKey *string         `json:"idempotency_key,omitempty"`
	LastError      *string         `json:"last_error,omitempty"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

type DeadJob struct {
	ID                uuid.UUID       `json:"id"`
	QueueName         string          `json:"queue_name"`
	Payload           json.RawMessage `json:"payload"`
	JobType           string          `json:"job_type"`
	Attempts          int             `json:"attempts"`
	LastError         *string         `json:"last_error,omitempty"`
	OriginalCreatedAt time.Time       `json:"original_created_at"`
	DiedAt            time.Time       `json:"died_at"`
}
