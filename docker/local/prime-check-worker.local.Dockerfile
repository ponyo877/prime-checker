FROM golang:1.24-alpine

# Install air for hot reload
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Create air config if not exists
RUN if [ ! -f .air.toml ]; then air init; fi

# Expose port for debugging if needed
EXPOSE 40002

CMD ["air", "-c", ".air.toml", "--", "cmd/prime-check-worker"]