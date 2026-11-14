```
  ________                  ____  __.            .__ 
 /  _____/  ____           |    |/ _|____ ______ |__|
/   \  ___ /  _ \   ______ |      < /  _ \\____ \|  |
\    \_\  (  <_> ) /_____/ |    |  (  <_> )  |_> >  |
 \______  /\____/          |____|__ \____/|   __/|__|
        \/                         \/     |__|       
```

A Go backend with clean architecture, supporting MySQL, Redis, NSQ, and Temporal workflows. Register and activation run through the onboarding workflow.

## Quick Start

### Ask Barista ☕

Create a new project by asking your barista to brew one:

```bash
# Brew a fresh project (latest stable tag: v3.1.0)
curl -fsSL https://raw.githubusercontent.com/RandySteven/go-kopi/v3.1.0/barista | bash -s -- brew -n my-project

# Or with your own git remote
curl -fsSL https://raw.githubusercontent.com/RandySteven/go-kopi/v3.1.0/barista | bash -s -- brew -n my-project -r https://github.com/youruser/my-project.git
```

### Manual Installation

```bash
# Clone the repository at the latest stable tag
git clone --branch v3.1.0 --single-branch https://github.com/RandySteven/go-kopi.git my-project
cd my-project

# Run setup
./barista setup

# Change remote to your own repository
./barista remote -r https://github.com/youruser/my-project.git
```

### Barista Commands

| Command | Description |
|---------|-------------|
| `./barista brew -n <name>` | Brew a fresh project (clone + setup + remote) |
| `./barista clone -n <name>` | Clone to a new project directory |
| `./barista setup` | Set up config files and install dependencies |
| `./barista deps` | Install go-cook pins and optional deps (`install_mod.sh`) |
| `./barista remote -r <url>` | Change git remote to your own repo |
| `./barista refill` | Merge latest upstream template (default tag: `v3.1.0`) |
| `./barista help` | Show help message |

### Keeping Up to Date

After changing your remote, you can still get refills from the original go-kopi:

```bash
./barista refill
```

This will merge the latest changes while preserving your customizations.

## Tech Stack

| Category        | Technology |
|-----------------|------------|
| Language        | Go 1.27 |
| HTTP Router     | [Gorilla Mux](https://github.com/gorilla/mux) |
| Database        | MySQL ([go-cook/db](https://github.com/RandySteven/go-cook)) |
| Document Store  | MongoDB (mongo-driver) |
| Cache           | Redis (go-redis/v9) |
| Message Queue   | NSQ |
| Workflows       | [Temporal](https://temporal.io/) |
| Authentication  | JWT (golang-jwt/v5) |
| Logging         | Logrus |
| Scheduler       | Cron (robfig/cron) |
| Infrastructure  | [go-cook](https://github.com/RandySteven/go-cook) (DB, Redis, NSQ, Temporal clients) |
| AWS             | [go-baker](https://github.com/RandySteven/go-baker) (S3, DynamoDB) |

## Architecture

This project follows a **layered architecture** pattern:

```
┌─────────────────────────────────────────────────────────────┐
│                        HTTP Request                         │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                     middlewares                             │
│         (logging, CORS, auth, rate limiting, timeout)       │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                         routes                              │
│              (Route registration, prefixes)                 │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        handlers                             │
│              (Request parsing, validation)                  │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        usecases                             │
│                   (Business logic)                          │
└─────────────────────────────────────────────────────────────┘
                              │
              ┌───────────────┼───────────────┐
              ▼               ▼               ▼
┌─────────────────┐  ┌─────────────┐  ┌───────────────┐
│   repositories  │  │   caches    │  │  NSQ/Temporal │
│    (MySQL)      │  │   (Redis)   │  │  (Publish)    │
└─────────────────┘  └─────────────┘  └───────────────┘
```

The `apps` package wires infrastructure clients and exposes factory methods (`PrepareHttpHandler`, `PrepareConsumer`, `PrepareJobScheduler`) used by each `cmd` entry point.

### Async Processing (Consumers)

```
┌─────────────────────────────────────────────────────────────┐
│                    NSQ Message Queue                        │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                       consumers                             │
│              (Message handlers by topic)                    │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        usecases                             │
│                   (Business logic)                          │
└─────────────────────────────────────────────────────────────┘
                              │
              ┌───────────────┴───────────────┐
              ▼                               ▼
┌─────────────────────┐            ┌─────────────────┐
│     repositories    │            │     caches      │
│       (MySQL)       │            │     (Redis)     │
└─────────────────────┘            └─────────────────┘
```

## Project Structure

| Directory    | Description |
|--------------|-------------|
| apps         | Application bootstrap — wires DB, cache, NSQ, Temporal, and layer factories |
| apperror     | Application error types |
| caches       | Redis cache layer |
| cmd          | Entry points: `main`, `migration`, `drop`, `scheduler` |
| configs      | YAML config loading, validation, and HTTP server runner |
| consumers    | NSQ message consumers and runner registration |
| docker       | Dockerfile and Docker Compose |
| entities     | Domain models and request/response payloads |
| enums        | Enums and constants (route prefixes, middleware keys) |
| files        | Configuration files (YAML, `.env`) |
| handlers     | HTTP handlers |
| interfaces   | Shared interface definitions |
| jobs         | Scheduled job handlers |
| logic        | Temporal workflows. `logic/onboarding` is register and activation |
| middlewares  | HTTP middleware (auth, logging, CORS, rate limiting) |
| proto        | gRPC/protobuf service definitions |
| queries      | Raw SQL queries and migrations |
| repositories | Data access layer |
| routes       | Route definitions and router initialization |
| topics       | NSQ topic definitions |
| usecases     | Business logic layer |
| utils        | Helper functions |

## Getting Started

### Prerequisites

- Go 1.27+
- MySQL
- Redis
- NSQ (optional, for async processing)
- Temporal (required for `POST /auth/register` and `POST /auth/activated`)

### Configuration

1. Copy environment file:

```bash
cp files/env/.env.example files/env/.env
```

2. Copy YAML config:

```bash
cp files/yaml/app.example.yml files/yaml/app.local.yml
```

3. Update `.env` with your environment:

```env
# prod | dev | test | staging
ENV=dev

SECRET_JWT='your-jwt-secret'
ACTIVATION_JWT_KEY='your-activation-secret'
```

`ACTIVATION_JWT_KEY` signs the onboarding activation token. When it is empty, the workflow uses a built-in development key.

4. Update `app.local.yml` with your database and server config:

```yaml
configs:
  server:
    host: "0.0.0.0"
    port: "8080"
    timeout:
      server: 30
      read: 15
      write: 10
      idle: 5

  db:
    db: "mysql"
    host: "localhost"
    port: "3306"
    dbname: "your_database"
    dbuser: "your_user"
    dbpass: "your_password"

  redis:
    host: "localhost"
    port: "6379"
    password: ""

  nsq:
    host: "localhost"
    port: "4150"
    topic: "your_topic"
    channel: "your_channel"
    maxInFlight: 10
    maxRequeueDelay: 60
    maxRequeueCount: 3

  temporal:
    host: "localhost"
    port: "7233"
    task_queue: "your_task_queue"
    namespace: "default"
```

### Make Commands

| Command           | Description |
|-------------------|-------------|
| `make run`        | Run the HTTP server (`cmd/main`) |
| `make migration`  | Run database migrations |
| `make seed`       | Seed the database |
| `make drop`       | Drop all database tables |
| `make refresh`    | Drop, migrate, and seed (full reset) |
| `make make_model` | Generate model and repository scaffold files |
| `make gen_proto`  | Generate Go code from `proto/service.proto` |
| `make run-docker` | Start with Docker Compose |
| `make stop-docker`| Stop Docker containers |

### Running the Application

```bash
# Set up database
make migration

# Run the server
make run
```

## Onboarding workflow

`POST /auth/register` starts Temporal workflow `onboarding_workflow_execution`. The handler returns after the user row exists and an activation token is ready. The workflow stays parked until `POST /auth/activated`.

| Step | Activity | What it does |
|------|----------|----------------|
| 1 | `persist_onboarding_request_activity` | Inserts `onboardings` with status `PENDING` |
| 2 | `register_user_activity` | Inserts `users` with status `PROCESSING`, sets onboarding to `ACTIVE`, and publishes the `RegisterResponse` query |
| 3 | `update_onboarding_status_activity` | Waits for signal `activated_user_signal`, then sets onboarding to `ONBOARDED` and the user to `ACTIVE` |

`register_as` defaults to `USER`. Any other value is rejected.

### `POST /auth/register`

```bash
curl -s -X POST http://localhost:8080/auth/register \
  -H 'Content-Type: application/json' \
  -d '{
    "first_name": "Randy",
    "last_name": "Steven",
    "email": "randy@example.com",
    "username": "randy",
    "password": "secret123",
    "phone_number": "08123456789",
    "address": "Jakarta",
    "register_as": "USER",
    "additional_info": { "occupation": "engineer" }
  }'
```

The `201` body includes `activation_token`. Copy it for activation. The token expires in 15 minutes.

### `POST /auth/activated`

```bash
curl -s -X POST http://localhost:8080/auth/activated \
  -H 'Content-Type: application/json' \
  -d '{"token":"<activation_token>"}'
```

A successful activation returns `200` with `"activate": true`.

`POST /auth/login` remains registered and is separate from this workflow.

### Testing

```bash
go test ./...
```

### Docker

```bash
# Start all services
make run-docker

# Stop all services
make stop-docker
```

## Adding New Features

### 1. Add a New HTTP Handler

1. Create handler in `handlers/` (e.g. `user_handler.go`)
2. Define the handler interface in the same package
3. Register the handler in `handlers/handler.go` via `NewHandlers`
4. Add routes in `routes/routers.go` using helpers from `routes/action.go` (`Get`, `Post`, `Put`, `Delete`, `Patch`)

Example route registration:

```go
endpointRouters[enums.AuthPrefix] = []*Router{
    Post("RegisterUser", "/register", api.OnboardingHandler.RegisterUser),
    Post("ActivationUser", "/activated", api.OnboardingHandler.ActivateUser),
    Post("LoginUser", "/login", api.UserHandler.LoginUser),
}
```

### 2. Add a New Consumer

1. Create consumer logic in `consumers/`
2. Define the topic in `topics/`
3. Register the consumer with `Runners.RegisterConsumer(topic, handler)`
4. Start consumers via `Runners.Run(ctx)`

### 3. Add a Scheduled Job

1. Create job handler in `jobs/`
2. Wire the job in `apps/app.go` → `PrepareJobScheduler`
3. Run via `make` target or `cmd/scheduler`

### 4. Add a New Model

```bash
make make_model
# Enter model name when prompted
```

This generates scaffold files for the model struct and repository interface.

## Environment Support

| Environment | Config File |
|-------------|-------------|
| dev         | `files/yaml/app.local.yml` |
| staging     | `files/yaml/app.docker.yml` |
| prod        | `files/yaml/app.prod.yml` |

Set environment in `.env`:

```env
ENV=dev
```

The Makefile selects the YAML file based on `ENV`.

---

<sub>This documentation was written entirely by AI because the developer was too lazy to write it themselves. You're welcome.</sub>
