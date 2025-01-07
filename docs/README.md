# Desafio Taghos Backend JR

## Overview
This project is a REST API built with Go, designed to manage items with basic CRUD operations. It serves as a backend service for applications that require item management functionalities.

## Project Structure
```
desafio_taghos_backend_jr
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── handler.go   # HTTP request handlers
│   ├── models
│   │   └── model.go     # Data structures
│   ├── routes
│   │   └── routes.go    # Application routes
│   └── services
│       └── service.go   # Business logic
├── pkg
│   └── utils
│       └── utils.go     # Utility functions
├── docs
│   └── README.md        # Project documentation
├── scripts
│   └── setup.sh         # Setup script
└── go.mod               # Module definition
```

## Setup Instructions
1. Clone the repository:
   ```
   git clone <repository-url>
   cd desafio_taghos_backend_jr
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go
   ```

## API Usage
- **GET /items**: Retrieve a list of items.
- **POST /items**: Create a new item. Requires a JSON body with `name` and `description`.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.