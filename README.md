# Go GORM REST API

> **Note:** This is a practice project for learning Go, GORM, and building REST APIs.

A simple REST API built with Go, GORM, and Gorilla Mux that demonstrates CRUD operations for Users and Tasks with PostgreSQL database.

## 🛠️ Tech Stack

- **Go** - Programming language
- **GORM** - ORM library for Go
- **Gorilla Mux** - HTTP router and URL matcher
- **PostgreSQL** - Database
- **Air** - Live reload for Go applications

## 📋 Prerequisites

Before running this project, make sure you have:

- Go 1.25.5 or higher installed
- PostgreSQL installed and running
- A PostgreSQL database named `gorm` (or update the connection string)

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone <your-repo-url>
cd go-gorm-rest-api
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Configure the database

Update the database connection string in `db/connection.go`:

```go
var DSN = "host=localhost user=your_user password=your_password dbname=gorm port=5432 sslmode=disable TimeZone=America/Caracas"
```

### 4. Run the application

**Option 1: Using Air (Recommended for development)**

First, make sure Go's bin directory is in your PATH. Add this to your `~/.zshrc` (or `~/.bashrc` for bash):

```bash
export PATH="$HOME/go/bin:$PATH"
```

Then reload your shell configuration:

```bash
source ~/.zshrc
```

Now install Air:

```bash
go install github.com/cosmtrek/air@latest
```

Then run:

```bash
air
```

**Option 2: Using Go directly**

```bash
go run main.go
```

The API will be available at `http://localhost:3000`

## 📡 API Endpoints

### Users

- `GET /` - Home endpoint
- `GET /users` - Get all users
- `GET /users/{id}` - Get a specific user
- `POST /users` - Create a new user
- `PUT /users/{id}` - Update a user
- `DELETE /users/{id}` - Delete a user (soft delete)

### Tasks

- `GET /tasks` - Get all tasks
- `GET /tasks/{id}` - Get a specific task
- `POST /tasks` - Create a new task
- `PUT /tasks/{id}` - Update a task
- `DELETE /tasks/{id}` - Delete a task (soft delete)

## 📝 Example Requests

### Create a User

```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com"
  }'
```

### Create a Task

```bash
curl -X POST http://localhost:3000/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Study Go programming language",
    "done": false,
    "user_id": 1
  }'
```

### Get All Users

```bash
curl http://localhost:3000/users
```

## 🗄️ Database Models

### User

- `ID` - Primary key
- `FirstName` - User's first name
- `LastName` - User's last name
- `Email` - User's email (unique)
- `Tasks` - Associated tasks (relationship)

### Task

- `ID` - Primary key
- `Title` - Task title (unique)
- `Description` - Task description
- `Done` - Completion status
- `UserID` - Foreign key to User

## 📁 Project Structure

```
go-gorm-rest-api/
├── db/
│   └── connection.go      # Database connection setup
├── models/
│   ├── Users.go           # User model
│   └── Tasks.go           # Task model
├── routes/
│   ├── index.routes.go    # Home route
│   ├── users.routes.go    # User routes
│   └── tasks.routes.go    # Task routes
├── main.go                 # Application entry point
└── go.mod                  # Go module dependencies
```

## 🎯 Learning Objectives

This project demonstrates:

- Setting up a Go REST API
- Database connection with GORM
- CRUD operations
- Model relationships (User-Task)
- HTTP routing with Gorilla Mux
- JSON request/response handling
- Error handling in Go

## 📄 License

This is a practice project for educational purposes.
