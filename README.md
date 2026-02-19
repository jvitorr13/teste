# Backend System Foundation

This is a scalable backend system foundation in Go using Clean Architecture and modular design.

## Project Structure

- `cmd/api`: Application entrypoint.
- `internal/{module}`: Domain, UseCases, Repositories, Handlers, and Infrastructure for each module.
- `pkg/`: Shared packages like config, database, and logger.

## Architecture Rules

- **Domain**: No external dependencies.
- **Usecases**: Depend only on interfaces.
- **Repositories**: Defined as interfaces in the repository layer.
- **Infrastructure**: Implements repositories and handles DB specifics.
- **Handlers**: Call usecases only and handle HTTP specifics.

## Modules

### User Module
- `POST /users`: Create a new user.
- `GET /users`: List all users.

### Task Module
- `POST /tasks`: Create a new task.
- `PUT /tasks/status?id={id}`: Update task status.

## Tech Stack
- Go
- standard `net/http` (Go 1.22+ routing features used)
- PostgreSQL
- sqlx

## Environment Variables
- `DB_URL`: PostgreSQL connection string.
- `SERVER_ADDR`: Address to run the HTTP server (default: `:8080`).
