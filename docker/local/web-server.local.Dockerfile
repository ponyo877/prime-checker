FROM golang:1.24-alpine

# Install air for hot reload
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Expose ports
EXPOSE 8080
EXPOSE 40003

CMD ["air", "-c", "./cmd/web-server/air.toml"]