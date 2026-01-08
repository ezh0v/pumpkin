# Development Guidelines

## Requirements

- [Golang (1.24+)](https://go.dev/)
- [Docker / Docker Compose](https://docs.docker.com/engine/)
- [Goose](https://github.com/pressly/goose) database migration tool

## Build Process

1. Install dependencies.

   ```sh
   go install github.com/pressly/goose/v3/cmd/goose@latest
   ```

2. Copy example environment file and update values inside as needed.

   ```sh
   cp .example.env .env
   ```

3. Start required services (Postgres, Minio, etc.):

   ```sh
   docker compose -f docker-compose.dev.yml up -d
   ```

4. Run application:

   ```sh
   go run main.go
   ```

5. Then open http://localhost:8000 in your browser.
