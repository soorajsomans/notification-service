# Notification Platform (Go + PostgreSQL)

A production-style Notification Platform built in Go.

This project demonstrates:

* REST APIs
* Clean Architecture
* Repository Pattern
* Worker Pattern
* PostgreSQL
* Retry Mechanism
* Exponential Backoff
* Notification Status Tracking

---

# Features

## API Layer

Create notifications via REST API.

```text
POST /api/v1/notifications
GET  /api/v1/notifications/:id
```

---

## Worker

Background worker continuously polls the database and processes notifications.

Flow:

```text
PENDING
   ↓
PROCESSING
   ↓
SENT
```

Failure Flow:

```text
PENDING
   ↓
PROCESSING
   ↓
RETRY
   ↓
PROCESSING
   ↓
RETRY
   ↓
FAILED
```

---

## Retry Support

Supports exponential backoff retries.

Example:

```text
Retry 1 → 2 seconds
Retry 2 → 4 seconds
Retry 3 → 8 seconds
Retry 4 → 16 seconds
Retry 5 → FAILED
```

---

# Project Structure

```text
notification-platform/

├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── database/
│   │   └── postgres.go
│   │
│   └── notification/
│       ├── handler/
│       │   └── notification_handler.go
│       │
│       ├── model/
│       │   └── notification.go
│       │
│       ├── provider/
│       │   └── email_provider.go
│       │
│       ├── repository/
│       │   └── postgres_repository.go
│       │
│       ├── service/
│       │   ├── notification_service.go
│       │   ├── retry.go
│       │   └── constants.go
│       │
│       └── worker/
│           └── worker.go
│
├── migrations/
│
├── go.mod
└── README.md
```

---

# Database Schema

## notifications

```sql
CREATE TABLE notifications (
    id UUID PRIMARY KEY,

    user_id VARCHAR(255) NOT NULL,

    channel VARCHAR(50) NOT NULL,

    message TEXT NOT NULL,

    status VARCHAR(50) NOT NULL,

    retry_count INT NOT NULL DEFAULT 0,

    next_retry_at TIMESTAMPTZ NULL,

    created_at TIMESTAMPTZ NOT NULL,

    updated_at TIMESTAMPTZ NOT NULL
);
```

---

# Notification Status

```text
PENDING
PROCESSING
RETRY
SENT
FAILED
```

---

# Run PostgreSQL

```bash
docker run \
  --name postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_DB=notification_db \
  -p 5432:5432 \
  -d postgres:16
```

---

# Create Table

Connect:

```bash
psql \
-h localhost \
-U postgres \
-d notification_db
```

Run:

```sql
CREATE TABLE notifications (
    id UUID PRIMARY KEY,

    user_id VARCHAR(255) NOT NULL,

    channel VARCHAR(50) NOT NULL,

    message TEXT NOT NULL,

    status VARCHAR(50) NOT NULL,

    retry_count INT NOT NULL DEFAULT 0,

    next_retry_at TIMESTAMPTZ NULL,

    created_at TIMESTAMPTZ NOT NULL,

    updated_at TIMESTAMPTZ NOT NULL
);
```

---

# Start Application

```bash
go run cmd/api/main.go
```

Expected:

```text
Connected to Postgres
Server running on :8080
Worker started
```

---

# API Examples

## Create Notification

```bash
curl --location 'http://localhost:8080/api/v1/notifications' \
--header 'Content-Type: application/json' \
--data '{
    "user_id":"user-123",
    "channel":"EMAIL",
    "message":"Welcome to Notification Platform"
}'
```

Response:

```json
{
"id": "9975125c-de28-4ef0-9768-38a7b10a227d"
}
```

---

## Get Notification

```bash
curl --location \
'http://localhost:8080/api/v1/notifications/NOTIFICATION_ID'
```

Response:

```json
{
  "id":"4db8dbd7-9d55-4d92-a2d1-4fdde4d8ab59",
  "user_id":"user-123",
  "channel":"EMAIL",
  "message":"Welcome to Notification Platform",
  "status":"SENT",
  "retry_count":0
}
```

---

# Retry Testing

Force provider failure:

```go
func (p *EmailProvider) Send(
    ctx context.Context,
    notification model.Notification,
) error {

    return errors.New("provider timeout")
}
```

Create a notification.

Worker log:

```text
provider is down retrying

Marking for retry

Successfully marked for retry
```

Database:

```sql
SELECT
    id,
    status,
    retry_count,
    next_retry_at
FROM notifications;
```

Example:

```text
RETRY | 1 | 2026-06-20 14:00:02+00
```

After retry limit:

```text
FAILED
```

---

# Current Architecture

```text
                ┌──────────────┐
                │ REST API     │
                └──────┬───────┘
                       │
                       ▼
                ┌──────────────┐
                │ PostgreSQL   │
                └──────┬───────┘
                       │
                       ▼
                ┌──────────────┐
                │ Worker       │
                └──────┬───────┘
                       │
                       ▼
                ┌──────────────┐
                │ Provider     │
                │ (Email)      │
                └──────────────┘
```

---

# Completed

* Notification API
* PostgreSQL Repository
* Worker Polling
* Claim Pattern
* Status Tracking
* Retry Mechanism
* Exponential Backoff
* Timezone-safe Retry Scheduling

---

# Next Steps

Planned enhancements:

* Dead Letter Queue (DLQ)
* Kafka Integration
* SMS Provider
* Push Notifications
* Provider Failover
* Outbox Pattern
* Metrics & Monitoring
* Prometheus
* OpenTelemetry Tracing
* Rate Limiting
* User Notification Preferences
* Scheduled Notifications
* Horizontal Worker Scaling

```
```
