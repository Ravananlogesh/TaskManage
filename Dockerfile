FROM golang:1.23 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build the Go application
RUN go build -o taskmanage main.go

# Final lightweight image
FROM alpine:latest
WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/taskmanage .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./taskmanage"]
