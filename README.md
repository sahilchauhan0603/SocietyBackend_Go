# Society Management Backend (Go)

REST API backend for the Society Management platform, organized for production-style deployment with modular routes, centralized configuration, middleware stack, and graceful shutdown.

## Production-Style Architecture

- cmd/api: application entrypoint
- internal/config: centralized env loading and validation
- internal/database: database bootstrap and migrations
- internal/http/middleware: request middleware
- internal/http/routes: split route registration by domain
- internal/helpers: reusable helper utilities (mail, OTP, reset)
- internal/models: GORM models
- internal/repository: repository interfaces and implementations by domain
- internal/service: business logic layer by domain
- internal/utils: shared utility code

## Directory Structure

```text
GoBACKEND/
  cmd/
    api/
      main.go
  internal/
    config/
      config.go
    database/
      database.go
    http/
      middleware/
      routes/
        routes.go
        auth_routes.go
        admin_routes.go
        society_routes.go
        student_routes.go
        content_routes.go
    helpers/
    models/
    repository/
      admin_repository.go
      society_repository.go
      student_repository.go
      content_repository.go
      health_repository.go
    service/
      admin/
      auth/
      content/
      society/
      student/
      health_service.go
    utils/
  .env
  .env.example
  .gitignore
  go.mod
  go.sum
  README.md
```

## Prerequisites

- Go 1.22.4 or newer
- PostgreSQL-compatible database
- SMTP credentials for email workflows
- Microsoft OAuth app credentials (if OAuth routes are used)

## Environment Variables

Copy `.env.example` to `.env` and update values.

Production-oriented server settings were added:

- `SERVER_READ_TIMEOUT`
- `SERVER_WRITE_TIMEOUT`
- `SERVER_IDLE_TIMEOUT`
- `SERVER_SHUTDOWN_TIMEOUT`
- `CORS_ALLOWED_ORIGINS` (comma-separated)
- `DB_SSLMODE`
- `DB_CONNECT_MAX_RETRIES`
- `DB_CONNECT_RETRY_INTERVAL`

Minimum template:

```env
# Server
PORT=8000
SERVER_READ_TIMEOUT=15s
SERVER_WRITE_TIMEOUT=15s
SERVER_IDLE_TIMEOUT=60s
SERVER_SHUTDOWN_TIMEOUT=10s
CORS_ALLOWED_ORIGINS=http://localhost:5173,http://localhost:8000,https://societymanagementfrontend-h3v3.onrender.com

# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=society_db
DB_USER=postgres
DB_PASS=your_db_password
DB_SSLMODE=verify-full
DB_CONNECT_MAX_RETRIES=12
DB_CONNECT_RETRY_INTERVAL=5s

# JWT
JWT_KEY=replace_with_secure_secret

# Default admin seed
ADMIN_USER=admin
ADMIN_PASS=admin_password

# SMTP
SMTP_SERVER=smtp.gmail.com
SMTP_PORT=587
EMAIL_USER=your_email@example.com
EMAIL_PASSWORD=your_email_app_password

# Microsoft OAuth
CLIENT_ID=your_microsoft_client_id
CLIENT_SECRET=your_microsoft_client_secret
TENANT_ID=your_tenant_id
REDIRECT_URL=http://localhost:8000/callback
```

## Run Commands

### Start server (new layout)

```powershell
cd "d:\PROJECTS - FULL_STACK\SocietyManagement\GoBACKEND"
go mod tidy
go run ./cmd/api
```

### Build binary

```powershell
cd "d:\PROJECTS - FULL_STACK\SocietyManagement\GoBACKEND"
go build -o society-backend.exe ./cmd/api
.\society-backend.exe
```

## Docker (Production-like)

Build and run API + PostgreSQL with healthchecks:

```powershell
cd "d:\PROJECTS - FULL_STACK\SocietyManagement\GoBACKEND"
docker compose -p society-backend up --build -d
```

View status:

```powershell
docker compose -p society-backend ps
docker compose -p society-backend logs -f api
```

Stop:

```powershell
docker compose -p society-backend down
```

Run again (after stop):

```powershell
docker compose -p society-backend up -d
```

Run again with rebuild (if code or Dockerfile changed):

```powershell
docker compose -p society-backend up --build -d
```

## API Base URL

- http://localhost:8000

## Health Endpoints

- GET /healthz
- GET /readyz

## Quick Smoke Test

```powershell
Invoke-RestMethod "http://localhost:8000/api/v1/news" -Method GET
Invoke-RestMethod "http://localhost:8000/healthz" -Method GET
```

## Notes

- Most routes are under `/api/v1`.
- OAuth and auth utility routes include `/microsoftLogin`, `/callback`, `/adminlogin`, `/forgotPassword`, and `/resetPassword`.
- Startup auto-runs DB migration and seeds default admin if missing.
- Graceful shutdown is enabled for SIGINT/SIGTERM.
- Middleware stack includes panic recovery, request ID propagation, structured JSON request logging, security headers, CORS, and trace ID injection in JSON responses.
- Request trace ID is returned via `X-Request-ID` header and injected as `trace_id` in JSON response payloads.
