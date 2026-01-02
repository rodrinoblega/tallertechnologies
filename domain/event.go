package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewEvent(title, description string, startTime, endTime time.Time) (*Event, error) {

	if title == "" {
		return nil, errors.New("title should not be empty")
	}

	if len(title) > 100 {
		return nil, errors.New("max title length: 100 characters")
	}

	if !startTime.Before(endTime) {
		return nil, errors.New("start time should be before end time")
	}

	return &Event{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		StartTime:   startTime,
		EndTime:     endTime,
		CreatedAt:   time.Now(),
	}, nil
}
