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
	ID uuid.UUID `db:"id" json:"id"`

	QueueName string `db:"queue_name" json:"queue_name"`

	Payload json.RawMessage `db:"payload" json:"payload"`

	JobType string `db:"job_type" json:"job_type"`

	Status JobStatus `db:"status" json:"status"`

	Priority int `db:"priority" json:"priority"`

	Attempts int `db:"attempts" json:"attempts"`

	MaxAttempts int `db:"max_attempts" json:"max_attempts"`

	RunAt time.Time `db:"run_at" json:"run_at"`

	LockedBy *string `db:"locked_by,omitempty" json:"locked_by,omitempty"`

	LockedUntil *time.Time `db:"locked_until,omitempty" json:"locked_until,omitempty"`

	IdempotencyKey *string `db:"idempotency_key,omitempty" json:"idempotency_key,omitempty"`

	LastError *string `db:"last_error,omitempty" json:"last_error,omitempty"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`

	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewJob(queueName, jobType string, payload json.RawMessage, priority, maxAttempts int) *Job {
	return &Job{
		ID:          uuid.New(),
		QueueName:   queueName,
		JobType:     jobType,
		Payload:     payload,
		Status:      StatusQueued,
		Priority:    priority,
		Attempts:    0,
		MaxAttempts: maxAttempts,
		RunAt:       time.Now().UTC(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
