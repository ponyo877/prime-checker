# Product Expiry Tracker

A microservice-based prime number checking system built with Go, using the Outbox pattern for reliable message processing.

## Architecture

The system consists of four main applications:

1. **Web Server** (`cmd/web-server`) - HTTP API server that receives prime check requests
2. **Outbox Publisher** (`cmd/outbox-publisher`) - Publishes messages from the outbox table to Redis Streams
3. **Prime Check Worker** (`cmd/prime-check-worker`) - Consumes prime check messages and performs calculations
4. **Email Send Worker** (`cmd/email-send-worker`) - Sends email notifications with prime check results

## Directory Structure

```
prime-checker/
├── cmd/                           # Application entry points
│   ├── web-server/               # HTTP API server
│   ├── outbox-publisher/         # Outbox pattern publisher
│   ├── prime-check-worker/       # Prime number calculation worker
│   └── email-send-worker/        # Email notification worker
├── internal/                      # Shared business logic
│   ├── adapter/                  # HTTP handlers
│   ├── model/                    # Domain models
│   ├── usecase/                  # Business logic
│   ├── repository/               # Data access layer
│   ├── message/                  # Message processing
│   │   ├── broker/               # Message broker interfaces
│   │   ├── publisher/            # Message publishers
│   │   └── consumer/             # Message consumers
│   ├── email/                    # Email sending
│   └── prime/                    # Prime number checking
├── db/                           # Database related files
│   ├── queries/                  # SQL queries
│   ├── generated_sql/            # Generated Go code from sqlc
│   └── init/                     # Database initialization
├── docker/                       # Docker configurations
├── deployments/                  # Deployment configurations
└── typespec/                     # API specification
```

## Technologies Used

- **Go** - Programming language
- **MySQL** - Database for storing prime check requests and outbox messages
- **NATS JetStream** - Message broker for inter-service communication
- **Mailpit** - Email testing tool for development
- **TypeSpec** - API specification
- **Docker** - Containerization
- **sqlc** - SQL code generation

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.24+ (for local development)

### Running with Docker

1. Start all services:
```bash
cd deployments
docker-compose up --build
```

2. Test the API:
```bash
curl -s -X POST http://localhost:8080/prime-check \
  -H "Content-Type: application/json" \
  -d '{"number": "17"}' | jq
```

3. List all prime check requests:
```bash
curl -s http://localhost:8080/prime-check | jq
```

### Running for Development

1. Start infrastructure services:
```bash
cd deployments
docker-compose up mysql nats mailpit
```

2. Run each application in separate terminals:
```bash
# Terminal 1: Web Server
go run cmd/web-server/main.go

# Terminal 2: Outbox Publisher
go run cmd/outbox-publisher/main.go

# Terminal 3: Prime Check Worker
go run cmd/prime-check-worker/main.go

# Terminal 4: Email Send Worker
go run cmd/email-send-worker/main.go
```

3. View sent emails:
   - Open http://localhost:8025 to access Mailpit web interface
   - All emails sent by the application will be captured and displayed here

## Environment Variables

### Database Configuration
- `MYSQL_HOST` - MySQL host (default: localhost)
- `MYSQL_PORT` - MySQL port (default: 3306)
- `MYSQL_DATABASE` - Database name
- `MYSQL_USER` - Database user
- `MYSQL_PASSWORD` - Database password

### NATS Configuration
- `NATS_URL` - NATS server URL (default: nats://localhost:4222)

### Email Configuration
- `SMTP_HOST` - SMTP server host (default: localhost for mailpit)
- `SMTP_PORT` - SMTP server port (default: 1025 for mailpit)
- `SMTP_USERNAME` - SMTP username (default: test@example.com)

## API Endpoints

### Prime Check
- `POST /prime-check` - Submit a number for prime checking
- `GET /prime-check` - List all prime check requests
- `GET /prime-check/{id}` - Get specific prime check request

### Settings
- `GET /settings` - Get application settings
- `POST /settings` - Update application settings

## Message Flow

1. Client sends prime check request to Web Server
2. Web Server stores request in database and creates message in outbox table
3. Outbox Publisher reads from outbox table and publishes to NATS JetStream
4. Prime Check Worker consumes message, performs calculation, and creates email message
5. Email Send Worker consumes email message and sends notification

## Database Schema

### Tables
- `users` - User information with auth tokens
- `prime_checks` - Prime check requests
- `outbox` - Outbox pattern messages for reliable delivery

## Development

### Code Generation

Generate Go code from SQL queries:
```bash
task sqlc
```

Generate OpenAPI code from TypeSpec:
```bash
cd typespec
npm install
npm run build
```

### Building Applications

Build all applications:
```bash
go build -o bin/web-server cmd/web-server/main.go
go build -o bin/outbox-publisher cmd/outbox-publisher/main.go
go build -o bin/prime-check-worker cmd/prime-check-worker/main.go
go build -o bin/email-send-worker cmd/email-send-worker/main.go
```

## Monitoring and Logging

All applications log to stdout with structured logging. Key events include:
- Prime check requests received
- Messages published/consumed
- Prime calculations completed
- Emails sent
- Errors and failures

## Scaling

Each component can be scaled independently:
- **Web Server**: Scale horizontally behind a load balancer
- **Outbox Publisher**: Single instance recommended to avoid duplicate processing
- **Prime Check Worker**: Scale horizontally for increased throughput
- **Email Send Worker**: Scale horizontally for high email volume

## Contributing

1. Follow Go best practices
2. Write tests for new functionality
3. Update documentation for API changes
4. Ensure all services can be built and run successfully