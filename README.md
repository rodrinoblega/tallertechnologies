# Go Events Service

## Overview

This project is a simple backend service in Go that manages a collection of Events. It follows Clean Architecture principles, separating the code into layers:

- Entities (Domain): business rules and validations.

- Use Cases: Business logic for creating, listing, and retrieving events.

- Adapters: concrete implementation of the repository and HTTP handlers.

- Frameworks: connection to PostgreSQL, external libraries.

Clear separation of dependencies: use cases depend on interfaces, not concrete implementations.

The service exposes RESTful endpoints to create, list, and fetch events by ID.

## SQL Schema

```
CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(100) NOT NULL,
    description TEXT,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
); 
```

## Running the Go service
From the root of the project, simply execute:

```
go run main.go
```

- The service will start at http://localhost:8080.

- The project uses a database hosted in Supabase.

- Go automatically uses the DSN in your main.go to connect to the database.

## Endpoints
| Method | Endpoint       | Description                              |
| ------ | -------------- | ---------------------------------------- |
| POST   | `/events`      | Create a new event                        |
| GET    | `/events`      | List all events (ordered by start_time)   |
| GET    | `/events/{id}` | Get a single event by UUID                |


## Example request

### Create Event

```
curl -X POST http://localhost:8080/events \
-H "Content-Type: application/json" \
-d '{
  "title": "Fútbol Training",
  "description": "Team practice",
  "start_time": "2026-01-03T10:00:00Z",
  "end_time": "2026-01-03T12:00:00Z"
}'
```

### List Events

```
curl http://localhost:8080/events 
```

### Get Event By ID

```curl http://localhost:8080/events/<uuid>```

## Notes

- Go service uses database/sql and lib/pq for PostgreSQL interaction.

- Input validation ensures:
    Title is non-empty and ≤ 100 characters.
    Start_time is before end_time.

- The service uses context for all DB queries for safety and timeout handling.

- Clean Architecture ensures business logic is decoupled from HTTP and DB, making the system maintainable and testable.

## Questions

* [rnoblega@gmail.com](rnoblega@gmail.com)