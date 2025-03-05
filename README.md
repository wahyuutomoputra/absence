# Absence Management System API

A simple attendance management system API built with Go and Gin framework.

## Features

- User management (register, login, update, delete)
- Attendance management (check-in, check-out)
- Authentication using JWT
- API Documentation with Swagger
- Hot reload for development

## Prerequisites

- Go 1.21 or higher
- MySQL
- Air (optional, for hot reload)

## Installation

1. Clone the repository
```bash
git clone https://github.com/wahyuutomoputra/absence
cd absence
```

2. Install dependencies
```bash
go mod download
```

3. Set up environment variables by copying the example file
```bash
cp .env.example .env
```

4. Update the `.env` file with your database credentials and other configurations

5. Install Air (optional, for hot reload)
```bash
go install github.com/air-verse/air@latest

# Add Go binary path to your shell configuration (if not already added)
echo 'export PATH=$HOME/go/bin:$PATH' >> ~/.zshrc
source ~/.zshrc  # Or restart your terminal
```

Note: The last two commands add the Go binary path to your system PATH. This is required to run tools like `air` and `swag` from anywhere in your terminal. If you're using a different shell (like bash), replace `.zshrc` with your shell's config file (e.g., `.bashrc`).

## Running the Application

### Normal Run
```bash
go run cmd/api/main.go
```

### Development Mode with Hot Reload
```bash
air
```

The application will start on `http://localhost:8080` by default.

## API Documentation

Swagger documentation is available at:
```
http://localhost:8080/swagger/index.html
```

To regenerate Swagger documentation after making changes:
```bash
swag init -g cmd/api/main.go
```

## API Endpoints

### Public Routes
- POST `/api/register` - Register new user
- POST `/api/login` - User login

### Protected Routes (Requires Authentication)
#### User Routes
- GET `/api/users/:id` - Get user by ID
- PUT `/api/users/:id` - Update user
- DELETE `/api/users/:id` - Delete user
- GET `/api/users/:id/attendance` - Get user attendance history

#### Attendance Routes
- POST `/api/attendance/check-in` - Record check-in
- POST `/api/attendance/check-out` - Record check-out
- GET `/api/attendance/:id` - Get attendance by ID

## Authentication

Protected routes require a Bearer token in the Authorization header:
```
Authorization: Bearer <your-token>
```

The token is obtained from the login endpoint response. 