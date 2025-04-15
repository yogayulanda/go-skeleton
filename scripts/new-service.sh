#!/bin/bash

SERVICE_NAME=$1
CAP_NAME="$(tr '[:lower:]' '[:upper:]' <<< ${SERVICE_NAME:0:1})${SERVICE_NAME:1}"

PROTO_PATH="api/proto/gen/v1/${SERVICE_NAME}.proto"
HANDLER_PATH="internal/handler/${SERVICE_NAME}_handler.go"
DOMAIN_PATH="internal/domain/${SERVICE_NAME}/service.go"

mkdir -p $(dirname $PROTO_PATH)
mkdir -p $(dirname $HANDLER_PATH)
mkdir -p $(dirname $DOMAIN_PATH)

# 1. Create proto stub
cat <<EOF > $PROTO_PATH
syntax = "proto3";

package v1;

option go_package = "github.com/yogayulanda/if-trx-history/api/proto/gen/v1;v1pb";

service ${CAP_NAME}Service {
  rpc DoSomething(${CAP_NAME}Request) returns (${CAP_NAME}Response);
}

message ${CAP_NAME}Request {
  string example = 1;
}

message ${CAP_NAME}Response {
  string result = 1;
}
EOF

# 2. Create handler
cat <<EOF > $HANDLER_PATH
package handler

import (
  "context"
  v1pb "github.com/yogayulanda/if-trx-history/api/proto/gen/v1"
)

type ${CAP_NAME}Handler struct {
  v1pb.Unimplemented${CAP_NAME}ServiceServer
}

func New${CAP_NAME}Handler() *${CAP_NAME}Handler {
  return &${CAP_NAME}Handler{}
}

func (h *${CAP_NAME}Handler) DoSomething(ctx context.Context, req *v1pb.${CAP_NAME}Request) (*v1pb.${CAP_NAME}Response, error) {
  return &v1pb.${CAP_NAME}Response{Result: "Hello from ${SERVICE_NAME}!"}, nil
}
EOF

# 3. Create domain placeholder
cat <<EOF > $DOMAIN_PATH
package ${SERVICE_NAME}

type Service interface {
  // define methods
}
EOF

echo "✅ Service $SERVICE_NAME generated!"
echo "ℹ️  Jangan lupa untuk:"
echo "- Tambahkan Register di protocol/grpc/server.go"
echo "- Tambahkan Register REST di grpc-gateway/server.go"
echo "- Tambahkan New di di/container.go"
