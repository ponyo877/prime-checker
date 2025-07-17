FROM golang:1.24-alpine

# Install air for hot reload
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Expose port for debugging if needed
EXPOSE 40001

CMD ["air", "-c", "./cmd/outbox-publisher/air.toml"]