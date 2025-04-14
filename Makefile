BUF_VERSION ?= 1.52.1
BUF_URL = https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-Linux-x86_64.tar.gz
INSTALL_DIR = /usr/local/bin
BUF_BIN = $(INSTALL_DIR)/buf

PROTOC_GEN_GRPC_GATEWAY_VERSION = v2.18.0
PROTOC_GEN_OPENAPIV2_VERSION = v2.18.0

# Direktori yang digunakan
PROTO_DIR = proto
GEN_DIR = gen/proto

.PHONY: install-buf install-plugins clean install setup generate all

install-buf:
	@echo "🔧 Installing Buf CLI v$(BUF_VERSION)..."

	@if [ -d "$(BUF_BIN)" ]; then \
		echo "⚠️  Found directory at $(BUF_BIN), removing..."; \
		sudo rm -rf $(BUF_BIN); \
	fi

	@curl -sSL $(BUF_URL) -o /tmp/buf.tar.gz
	@tar -xzf /tmp/buf.tar.gz -C /tmp
	@sudo mv /tmp/buf/bin/buf $(BUF_BIN)
	@sudo chmod +x $(BUF_BIN)

	@echo "✅ Buf installed at $(BUF_BIN)"
	@$(BUF_BIN) --version

install-plugins:
	@echo "🔌 Installing gRPC Gateway and OpenAPI plugins..."
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@$(PROTOC_GEN_GRPC_GATEWAY_VERSION)
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@$(PROTOC_GEN_OPENAPIV2_VERSION)

generate:
	@echo "⚙️  Generating protobuf files with Buf..."
	buf generate

clean:
	@echo "🧹 Cleaning up generated files..."
	rm -rf $(GEN_DIR)
	rm -rf gen/swagger

install: install-buf install-plugins

setup: install generate

all: setup

# Run the server
.PHONY: start
start:
	go run main.go server
