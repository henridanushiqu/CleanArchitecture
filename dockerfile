# Build Stage (using Alpine with Go)
FROM golang:1.24-alpine AS builder

# Install required build dependencies
RUN apk add --no-cache git build-base

WORKDIR /app

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go binary (disable CGO for static binary)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp .

# Final Minimal Image
FROM alpine:latest

WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder /app/myapp .

# Expose port (adjust if different in code)
EXPOSE 8080

# Run app binary
CMD ["./myapp"]
