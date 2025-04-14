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
PROTO_VERSIONS = v1 v2

generate-proto:
	@echo "🚀 Generating proto files..."
	@for version in $(PROTO_VERSIONS); do \
		make generate-proto-version VERSION=$$version; \
	done

generate-proto-version:
	@echo "Generating proto files for version $(VERSION)..."
	@echo "Checking if version folder api/proto/$(VERSION) exists..."

	# Check if the folder for the specified version exists
	@if [ ! -d "api/proto/$(VERSION)" ]; then \
		echo "Error: Folder api/proto/$(VERSION) does not exist."; \
		exit 1; \
	fi

	@echo "Checking if gen folder for version $(VERSION) exists..."
	# Check if the gen folder for the version exists, if not, create it
	@if [ ! -d "api/proto/gen/$(VERSION)" ]; then \
		echo "Creating folder api/proto/gen/$(VERSION)..."; \
		mkdir -p api/proto/gen/$(VERSION); \
	fi

	@echo "Running protoc for Go and gRPC..."
	@protoc -I=api/proto -I=api/proto/google/api --go_out=api/proto/gen/$(VERSION) --go-grpc_out=api/proto/gen/$(VERSION) api/proto/$(VERSION)/*.proto || { echo 'Error: Go and gRPC code generation failed.'; exit 1; }
	@echo "Go and gRPC code generation completed for version $(VERSION)."

	@echo "Running protoc for gRPC Gateway..."
	@protoc -I=api/proto -I=api/proto/google/api --grpc-gateway_out=api/proto/gen/$(VERSION) api/proto/$(VERSION)/*.proto || { echo 'Error: gRPC Gateway code generation failed.'; exit 1; }
	@echo "gRPC Gateway code generation completed for version $(VERSION)."
	
	@echo "Running protoc for OpenAPI specification..."
	@mkdir -p api/swagger/$(VERSION)
	@protoc -I=api/proto -I=api/proto/google/api \
		--openapiv2_out=api/swagger/$(VERSION) \
		api/proto/$(VERSION)/*.proto || { echo 'Error: OpenAPI specification generation failed.'; exit 1; }

	# Move files from nested api/swagger/$(VERSION)/$(VERSION)/*.swagger.json to api/swagger/$(VERSION)
	@find api/swagger/$(VERSION)/$(VERSION) -name "*.swagger.json" -exec mv {} api/swagger/$(VERSION)/ \;
	@rm -rf api/swagger/$(VERSION)/$(VERSION)

	@echo "OpenAPI specification generation completed for version $(VERSION)."


# Variables
GO := go

# Run the server
.PHONY: start
start:
	$(GO) run main.go server

.PHONY: new-service
new-service:
	bash scripts/new-service.sh $(name)

.PHONY: create-service
create-service:
	@read -p "Enter service name (e.g. report): " service; \
	go run gen/gen.go $$service && make generate-proto
