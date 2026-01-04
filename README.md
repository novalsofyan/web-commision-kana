# Tech Stack
## Backend
- GoLang
- PostgreSQL
- sqlc
- goose

## Frontend (Planning)
- Vue.js or Nuxt

## DevOps
- Docker
- Docker Compose
- Kubernetes (planning to learn kubernetes)

# Run Project

Configuration on docker-compose.yaml & Dockerfile in projects folder

```bash
docker compose up -d
```

# Run Project for Dev
## Backend
- Create dotenv file & create key
  - APP_PORT
  - PROJECT_ENV
  - GOOSE_DBSTRING="host=localhost port=port user=user password=pass dbname=dbname sslmode=disable"
  - GOOSE_DRIVER=postgres
  - GOOSE_MIGRATION_DIR="./internal/repo/migrations"
  - USER_PASS_SETTING=pass
```bash
go mod run cmd/api/*.go
```
## Frontend
Cooming Soon!

## Docker Compose
- Create .docker.env file
 - POSTGRES_USER
 - POSTGRES_PASSWORD
 - POSTGRES_DB

 - DB_URL
 - APP_PORT
 - GOOSE_DBSTRING
 - GOOSE_DRIVER
 - GOOSE_MIGRATION_DIR
 - PROJECT_ENV
 - ALLOWED_ORIGINS
 