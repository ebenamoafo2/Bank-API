# Bank-API
A complete industry-ready JSON API project in Golang with JWT authentication, PostgreSql, and Docker.

## How to run with Docker

1. Make sure you have Docker and Docker Compose installed.
2. Run the following command in the root directory:
   ```bash
   docker compose up --build
   ```
3. The API will be available at `http://localhost:8080`.

## Endpoints

- `GET /account` - Get all accounts
- `GET /account/{id}` - Get account by ID
- `POST /account` - Create a new account
- `DELETE /account/{id}` - Delete an account
