# Variables
APP_NAME=my-go-app
GO_CMD=go
BUILD_DIR=bin
MAIN_FILE=cmd/api/main.go
TEST_DIR=./tests
MIGRATIONS_DIR=migrations

# Run the API server
run:
	$(GO_CMD) run $(MAIN_FILE)

# Build the Go application
build:
	$(GO_CMD) build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# Run tests
test:
	$(GO_CMD) test $(TEST_DIR)/... -v

# Lint the code (requires golangci-lint)
lint:
	golangci-lint run

# Format the code
fmt:
	$(GO_CMD) fmt ./...

# Apply database migrations (requires migrate tool)
migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "postgres://user:pass@localhost:5432/mydb?sslmode=disable" up

# Rollback the last migration
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "postgres://user:pass@localhost:5432/mydb?sslmode=disable" down 1

# Build and run with Docker
docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run --rm -p 8080:8080 --env-file .env $(APP_NAME)

# Clean the build directory
clean:
	rm -rf $(BUILD_DIR)/*
