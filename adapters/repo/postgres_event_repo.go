package repo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/rodrinoblega/tallertechnologies/domain"
)

type PostgresEventRepository struct {
	db *sql.DB
}

func NewPostgresEventRepository(db *sql.DB) *PostgresEventRepository {
	return &PostgresEventRepository{db: db}
}

func (r *PostgresEventRepository) Create(ctx context.Context, e *domain.Event) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO events (id, title, description, start_time, end_time, created_at)
         VALUES ($1,$2,$3,$4,$5,$6)`,
		e.ID, e.Title, e.Description, e.StartTime, e.EndTime, e.CreatedAt)
	return err
}

func (r *PostgresEventRepository) ListEvents(ctx context.Context) ([]*domain.Event, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, title, description, start_time, end_time, created_at 
         FROM events ORDER BY start_time ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*domain.Event
	for rows.Next() {
		e := &domain.Event{}
		if err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.StartTime, &e.EndTime, &e.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func (r *PostgresEventRepository) GetByID(ctx context.Context, id string) (*domain.Event, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	e := &domain.Event{}
	err = r.db.QueryRowContext(ctx,
		`SELECT id, title, description, start_time, end_time, created_at 
         FROM events WHERE id=$1`, uid).
		Scan(&e.ID, &e.Title, &e.Description, &e.StartTime, &e.EndTime, &e.CreatedAt)
	if err != nil {
		return nil, err
	}

	return e, nil
}
