# Dockerfile

# Build stage
FROM golang:1.23 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build the Go application
RUN go build -o taskmanage main.go

# Final image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Install necessary packages (optional, e.g., for TLS or debugging)
RUN apk --no-cache add ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /app/taskmanage .

# Copy configuration files (if needed)
COPY toml/config.toml /root/config.toml

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./taskmanage"]
