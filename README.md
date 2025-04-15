# go-skeleton

`go-skeleton` is a gRPC-based microservice designed to Service GO. It provides both gRPC and RESTful APIs, supports health checks, and integrates with SQL Server, Redis, and Kafka.

---

## Features

- **gRPC and REST API**: Offers gRPC endpoints and RESTful APIs via gRPC-Gateway.
- **Health Checks**: Monitors the health of dependencies like Redis, Kafka, and SQL Server.
- **Protobuf Integration**: Uses Protocol Buffers for API definitions.
- **Database Support**: SQL Server integration using GORM.
- **Logging**: Structured logging with `zap`.
- **Code Generation**: Automated generation of handlers, domains, and containers from `.proto` files.

---

## Project Structure

### Root Files
- **`go.mod`**: Go module dependencies.
- **`Makefile`**: Automates tasks like building, testing, and generating code.
- **`buf.yaml`**: Configuration for Buf, a tool for managing Protobuf files.
- **`buf.gen.yaml`**: Defines plugins for generating gRPC, REST, and OpenAPI code.

### Key Directories
- **`cmd/server`**: Contains the entry point for the gRPC server.
- **`gen`**: Auto-generated code from Protobuf files.
  - **`proto/v1`**: Generated gRPC and REST code for version 1 APIs.
  - **`proto/v2`**: Generated gRPC and REST code for version 2 APIs.
  - **`swagger`**: OpenAPI specifications for REST APIs.
- **`internal`**: Core business logic and utilities.
  - **`config`**: Application configuration management.
  - **`database`**: Database connection setup.
  - **`di`**: Dependency injection container.
  - **`domain`**: Business logic for different modules (e.g., `health`, `history`, `report`, `user`).
  - **`handler`**: gRPC and REST handlers for services.
  - **`middleware`**: Middleware for logging and error recovery.
  - **`protocol`**: Protocol-specific server configurations (gRPC and gRPC-Gateway).
  - **`utils`**: Utility functions (e.g., Swagger integration).
- **`proto`**: Source Protobuf files for API definitions.
  - **`v1`**: Protobuf files for version 1 APIs.
  - **`v2`**: Protobuf files for version 2 APIs.
- **`scripts`**: Helper scripts for creating new services.

---

## Prerequisites

- **Go**: Version 1.20 or later.
- **Protobuf Compiler**: Version 3.21.12 or later.
- **Buf CLI**: For managing Protobuf files.
- **SQL Server**: For database storage.
- **Redis**: For caching and health checks.
- **Kafka**: For message streaming.

---

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd go-skeleton
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Install Buf CLI:
   ```bash
   make install-buf
   ```

4. Generate Protobuf files:
   ```bash
   make generate
   ```

---

## Configuration

1. Copy the `.env` file and update the configuration:
   ```bash
   cp .env.example .env
   ```

2. Update the `.env` file with your environment-specific values:
   ```properties
   GRPC_PORT=8081
   HTTP_PORT=8080
   MSSQL_DB=mydb
   MSSQL_USER=root
   MSSQL_PASSWORD=password
   MSSQL_HOST=localhost
   MSSQL_PORT=1433
   ```

---

## Usage

### Run the Server

Start the server using the `Makefile`:
```bash
make start
```

Alternatively, run the server directly:
```bash
go run main.go server
```

### Run Tests

Run all tests:
```bash
make test
```

---

## Development

### Code Generation

Generate gRPC, REST, and OpenAPI code from Protobuf files:
```bash
make generate
```

### Debugging

Run the application in debug mode using Delve:
```bash
make debug
```

### Code Formatting

Format the codebase:
```bash
make fmt
```

### Linting

Run lint checks:
```bash
make lint
```

---

## Deployment

1. Build the application:
   ```bash
   make build
   ```

2. Deploy the binary (`go-skeleton`) to your server.

---

## API Documentation

### gRPC Endpoints

- **Health Check**: `/v1/health`
- **Get Transactions**: `/v1/transactions`
- **Store Transaction**: `/v1/transactions`

### REST Endpoints

- **Health Check**: `GET /v1/health`
- **Get Transactions**: `GET /v1/transactions`
- **Store Transaction**: `POST /v1/transactions`

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your feature"
   ```
4. Push to your branch:
   ```bash
   git push origin feature/your-feature-name
   ```
5. Open a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [gRPC](https://grpc.io/)
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [Zap Logging](https://github.com/uber-go/zap)
