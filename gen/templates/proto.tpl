syntax = "proto3";

package v1;

option go_package = "/.;v1";

service {{.Entity}}Service {
  rpc Sample (SampleRequest) returns (SampleResponse);
}

message SampleRequest {
  string id = 1;
}

message SampleResponse {
  string message = 1;
}
