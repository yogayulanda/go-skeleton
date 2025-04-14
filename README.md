# if-trx-history

`if-trx-history` is a gRPC-based microservice for managing transaction history. It includes a gRPC-Gateway for RESTful API access and supports health checks, transaction management, and integration with Redis, Kafka, and SQL databases.

---

## Features

- **gRPC and REST API**: Provides both gRPC and RESTful endpoints for managing transactions.
- **Health Checks**: Monitors the health of Redis, Kafka, and the database.
- **Protobuf Integration**: Uses Protocol Buffers for API definitions.
- **Environment Configuration**: Configurable via `.env` file.
- **Logging**: Structured logging using `zap`.
- **Database Support**: Supports MySQL, SQL Server, and MongoDB.

---

## Prerequisites

- **Go**: Version 1.20 or later.
- **Protobuf Compiler**: Version 3.21.12 or later.
- **Redis**: For caching and health checks.
- **Kafka**: For message streaming.
- **MySQL/SQL Server/MongoDB**: For database storage.

---

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd if-trx-history
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the environment:
   ```bash
   make setup-env
   ```

4. Generate Protobuf files:
   ```bash
   make generate-proto
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
   REDIS_ADDR=localhost:6379
   KAFKA_BROKER=localhost:9092
   DB_MYSQL_USER_APPS=root
   DB_MYSQL_PASS_APPS=password
   DB_MYSQL_NAME_APPS=mydb
   DB_MYSQL_HOST_APPS=localhost
   DB_MYSQL_PORT_APPS=3306
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

## Development

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

2. Deploy the binary (`if-trx-history`) to your server.

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




proto -> gen/proto/v1/user.pb.go      ✅ generated
         ↓
internal/handler/user_handler.go      ✅ interface ke proto
         ↓
internal/domain/user/service.go       ✅ logic bisnis
         ↓
internal/domain/user/repository.go    ✅ interface data access
         ↓
internal/database/user_repository.go  ✅ implementasi nyata (e.g., SQL)
