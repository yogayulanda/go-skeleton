# Variabel untuk versi dan lokasi binary
PROTOC_VERSION=3.21.12
PROTOC_URL=https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip
GOBIN=$(shell go env GOPATH)/bin

# Target untuk menginstal protoc dan plugin
setup-env:
	@echo " > Setting up environment"
	@echo " > Checking for protoc..."
	@if [ -x "$(shell which protoc)" ]; then \
		echo " > protoc already installed"; \
	else \
		echo " > protoc not found"; \
	fi
	@if [ -x "$(shell which protoc)" ]; then \
		echo " > protoc already installed"; \
	else \
		echo " > Installing protoc..."; \
		wget -q -O /tmp/protoc.zip $(PROTOC_URL); \
		unzip -o /tmp/protoc.zip -d /usr/local; \
		rm /tmp/protoc.zip; \
	fi
	@echo " > Installing Go plugins..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@echo " > Setup complete"
	
# Target untuk generate proto files
## generate-proto: Parse PROTO_FILES and generate output based on the options given
generate-proto:
	@echo " > Generate proto files"
	@mkdir -p api/proto/gen
	@mkdir -p api/proto/swagger
	protoc --proto_path=api/proto \
	 --go_out=api/proto/gen \
     --go_opt=paths=source_relative \
     --go-grpc_out=api/proto/gen \
     --go-grpc_opt=require_unimplemented_servers=false \
     --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=logtostderr=true:api/proto/gen \
	--openapiv2_out=logtostderr=true:api/proto/swagger \
	api/proto/*.proto
	@echo " > Generate proto files complete"

# Variables
GO := go

# Run the server
.PHONY: start
start:
	$(GO) run main.go server