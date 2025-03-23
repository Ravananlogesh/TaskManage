# Task Management API

This is a **Task Management API** built using **Go (Gin framework)**, providing user authentication, task management, rate limiting, and database interactions using PostgreSQL.

## Features

- **User Authentication (JWT)**: Register and login functionality
- **Task CRUD Operations**: Create, Read, Update, and Delete tasks
- **Rate Limiting**: Limits API requests to prevent abuse
- **IP Restriction**: Limits registration access based on IP
- **Database Integration**: Uses PostgreSQL with GORM
- **Logging**: Saves logs into a file with timestamps

## Installation & Setup

### Prerequisites

Ensure you have the following installed:

- Go (1.18+)
- PostgreSQL
- Git

### Clone the Repository

```sh
git clone https://github.com/your-repo/tasks-api.git
cd tasks-api
```

### Install Dependencies

```sh
go mod tidy
```

### Configure the Database

Update the `toml/config.toml` file with your database credentials:

```toml
[database]
host = "localhost"
port = 5432
user = "your_user"
password = "your_password"
dbname = "tasks_db"
sslmode = "disable"
```

### Run Database Migrations

```sh
go run migrations/migrate.go
```

### Start the Server

```sh
go run main.go
```

## API Endpoints

### User Authentication

| Method | Endpoint    | Description           |
| ------ | ----------- | --------------------- |
| POST   | `/register` | Register a new user   |
| POST   | `/login`    | Login & get JWT token |

### Task Management

| Method | Endpoint     | Description               |
| ------ | ------------ | ------------------------- |
| POST   | `/tasks/`    | Create a new task         |
| GET    | `/tasks/`    | Get all tasks (paginated) |
| GET    | `/tasks/:id` | Get task by ID            |
| PUT    | `/tasks/:id` | Update task               |
| DELETE | `/tasks/:id` | Delete task               |

### Middleware Used

- `AuthMiddleware()` → Protects task routes with JWT authentication
- `RateLimitMiddleware()` → Restricts excessive API calls
- `IPRestrictionMiddleware()` → Limits registration to specific IPs

## Logging

API logs are stored in `./log/` directory with a filename format:

```
logfileDDMMYYYY.HH.MM.SS.000000000.log
```

## Deployment

To run in production, build the application:

```sh
go build -o taskmanages main.go
./taskmanages
```
