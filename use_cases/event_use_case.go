package usecases

import (
	"context"
	"errors"

	"time"

	"github.com/rodrinoblega/tallertechnologies/domain"
)

type EventRepository interface {
	Create(ctx context.Context, e *domain.Event) error
	ListEvents(ctx context.Context) ([]*domain.Event, error)
	GetByID(ctx context.Context, id string) (*domain.Event, error)
}

type EventUseCase struct {
	repo EventRepository
}

func NewEventUseCase(repo EventRepository) *EventUseCase {
	return &EventUseCase{repo: repo}
}

func (uc *EventUseCase) CreateEvent(ctx context.Context, title, description string, startTime, endTime time.Time) (*domain.Event, error) {
	event, err := domain.NewEvent(title, description, startTime, endTime)

	if err != nil {
		return nil, err
	}

	err = uc.repo.Create(ctx, event)

	if err != nil {
		return nil, err
	}

	return event, nil

}

func (uc *EventUseCase) ListEvents(ctx context.Context) ([]*domain.Event, error) {
	return uc.repo.ListEvents(ctx)
}

func (uc *EventUseCase) GetEventByID(ctx context.Context, id string) (*domain.Event, error) {
	e, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("event not found")
	}

	return e, nil
}
