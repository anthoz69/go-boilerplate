# Go boilerplate API

This project is a Go-based API for boilerplate. It uses the Fiber web framework and follows a clean architecture pattern.

## Project Structure

```
.
├── cmd
│   └── app
│       └── main.go           # Application entry point
├── configs
│   └── config.yaml           # Configuration file
├── internal
│   ├── app
│   │   └── module.go         # Application module definitions
│   ├── configs
│   │   ├── config.go         # Configuration loading
│   │   ├── database.go       # Database configuration
│   │   └── redis.go          # Redis configuration
│   ├── core
│   │   ├── domain
│   │   │   └── user.go       # User domain model
│   │   └── ports
│   │       ├── repositories.go # Repository interfaces
│   │       └── usecases.go     # Use case interfaces
│   ├── modules
│   │   └── user
│   │       ├── user_handler.go    # HTTP handler for user-related endpoints
│   │       ├── user_module.go     # User module definition
│   │       ├── user_repository.go # User repository implementation
│   │       ├── user_router.go     # User route definitions
│   │       ├── user_usecase.go    # User use case implementation
│   │       └── user_usecase_test.go # Tests for user use case
│   └── router
│       └── router.go         # Main router setup
├── Makefile                  # Build and test commands
├── go.mod                    # Go module file
└── go.sum                    # Go module checksums
```

## Key Components

- `cmd/app/main.go`: The entry point of the application.
- `internal/app/module.go`: Defines the application modules using the `fx` dependency injection framework.
- `internal/core`: Contains the domain models and interfaces for the application.
- `internal/modules`: Contains feature-specific implementations (e.g., user module).
- `internal/router`: Sets up the main router for the application.
- `configs`: Contains configuration-related code and files.

## Running the Application

To run the application, use the following command:

```
go run cmd/app/main.go
```

## Running Tests

To run tests, use the following command:

```
make test
```

To generate and view a coverage report, use:

```
make coverage
```

This project uses Go modules for dependency management and follows a clean architecture pattern for better separation of concerns and testability.
