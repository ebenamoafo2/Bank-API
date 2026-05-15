# Bank-API

A complete industry-ready JSON API project in Golang featuring PostgreSQL and Docker.

## Overview
This project provides a simple banking API to manage user accounts. It supports creating, retrieving, and deleting accounts. The application is built with Go 1.26 and uses PostgreSQL for persistent storage.

## Requirements
- Go 1.26 or higher
- Docker and Docker Compose (recommended)
- PostgreSQL (if running locally without Docker)

## Setup & Run

### Using Docker (Recommended)
1. Make sure you have Docker and Docker Compose installed.
2. Run the following command in the root directory:
   ```bash
   docker compose up --build
   ```
3. The API will be available at `http://localhost:8080`.

### Running Locally
1. Install dependencies:
   ```bash
   go mod download
   ```
2. Create a `.env` file in the root directory and set your `DATABASE_URL`:
   ```env
   DATABASE_URL=postgres://user:password@localhost:5432/bankapi?sslmode=disable
   ```
3. Build and run the application:
   ```bash
   go run .
   ```
   The server will run on `http://localhost:3000` (default port in code when running locally).

## Scripts
- `go run .`: Start the application.
- `go build -o bin/bank-api`: Build the binary.
- `docker compose up`: Start the database and application in Docker.

## Environment Variables
The application uses the following environment variables:
- `DATABASE_URL`: Connection string for PostgreSQL (e.g., `postgres://user:password@host:port/dbname?sslmode=disable`).
- `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_PORT`: Used in `docker-compose.yml` for the database container.

## API Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/account` | Get all accounts |
| GET | `/account/{id}` | Get account by ID |
| POST | `/account` | Create a new account |
| DELETE | `/account/{id}` | Delete an account |

### Example Create Account Request
```json
POST /account
{
    "first_name": "John",
    "last_name": "Doe"
}
```

## Project Structure
- `main.go`: Application entry point.
- `api.go`: API server implementation and routing.
- `storage.go`: Database operations and PostgreSQL implementation.
- `types.go`: Core data structures and models.
- `Dockerfile`: Containerization instructions.
- `docker-compose.yml`: Multi-container setup for app and database.

## Tests
- [ ] TODO: Implement unit and integration tests.

## License
- [ ] TODO: Add license information.
