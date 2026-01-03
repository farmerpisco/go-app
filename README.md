# GitHub Actions Tutorial - Go Application

A simple Go web API for demonstrating GitHub Actions CI/CD workflows.

## Features

- Simple HTTP API with JSON responses
- Health check endpoint
- Greeting endpoint with query parameters
- Comprehensive test coverage
- Ready for GitHub Actions integration

## API Endpoints

- `GET /` - Root endpoint with welcome message
- `GET /health` - Health check endpoint
- `GET /api/greet?name=YourName` - Personalized greeting

## Local Development

### Prerequisites

- Go 1.21 or higher

### Running the Application

```bash
go run main.go
```

The server will start on port 8080 (or the PORT environment variable if set).

### Running Tests

```bash
go test -v
```

### Running with Coverage

```bash
go test -v -race -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Building

```bash
go build -o app .
./app
```

## Testing the API

```bash
# Root endpoint
curl http://localhost:8080/

# Health check
curl http://localhost:8080/health

# Greeting with name
curl http://localhost:8080/api/greet?name=Alice

# Greeting without name (defaults to "World")
curl http://localhost:8080/api/greet
```

## GitHub Actions

This project includes a complete CI/CD workflow (`.github/workflows/ci.yml`) that:

1. **Tests** - Runs all tests with race detection and coverage
2. **Build** - Compiles the application and uploads artifact

The workflow triggers on:
- Push to main branch
- Pull requests to main branch

## Project Structure

```
.
├── main.go                      # Application code
├── main_test.go                 # Test suite
├── go.mod                       # Go module definition
├── .github/
│   └── workflows/
│       └── go.yml              # GitHub Actions workflow
└── README.md                    # This file
```

## Using for Tutorials

This application is designed to be simple yet complete enough to demonstrate:
- Basic GitHub Actions setup
- Multiple job workflows
- Artifact uploading
- Code coverage reporting
- Testing with Go

Feel free to fork and modify for your own tutorials!

## Author: Paul Adegoke
