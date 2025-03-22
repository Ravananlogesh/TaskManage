# Use official Golang image
FROM golang:1.21 AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build the Go application
RUN go build -o my-go-app ./cmd/api/main.go

# Final lightweight image
FROM alpine:latest
WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/my-go-app .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./my-go-app"]
