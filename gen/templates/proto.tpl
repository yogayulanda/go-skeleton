syntax = "proto3";

package proto.v1;

option go_package = "/.;v1";

import "google/api/annotations.proto";

service {{.Entity}}Service {
  rpc HealthService (HealthRequest) returns (HealthResponse) {
    option (google.api.http) = {
      get: "/v1/{{.EntityLower}}/health"
    };
  }
}

message HealthRequest {
  string message = 1;
}

message HealthResponse {
  string health_status = 1; // e.g., "SERVING", "NOT_SERVING"
  map<string, string> component_statuses = 2; // e.g., { "db": "OK", "kafka": "FAIL", "redis": "OK" }
}
