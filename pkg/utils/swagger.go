package utils

import (
	"github.com/yogayulanda/go-skeleton/pkg/logging"
	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func LogAvailableEndpoints() {
	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		for i := 0; i < fd.Services().Len(); i++ {
			svc := fd.Services().Get(i)
			for j := 0; j < svc.Methods().Len(); j++ {
				m := svc.Methods().Get(j)
				opts := m.Options()
				if opts != nil {
					if httpRule, ok := proto.GetExtension(opts, annotations.E_Http).(*annotations.HttpRule); ok {
						logHttpRule(logging.Log, svc.FullName(), m.Name(), httpRule)
					}
				}
			}
		}
		return true
	})

	logging.Log.Info("✅ All available gRPC-Gateway endpoints have been listed")
}

func logHttpRule(logger *zap.Logger, serviceName protoreflect.FullName, methodName protoreflect.Name, rule *annotations.HttpRule) {
	var method, path string
	switch {
	case rule.GetGet() != "":
		method, path = "GET", rule.GetGet()
	case rule.GetPost() != "":
		method, path = "POST", rule.GetPost()
	case rule.GetPut() != "":
		method, path = "PUT", rule.GetPut()
	case rule.GetDelete() != "":
		method, path = "DELETE", rule.GetDelete()
	case rule.GetPatch() != "":
		method, path = "PATCH", rule.GetPatch()
	}

	logging.Log.Info("REST Endpoint registered",
		zap.String("method", method),
		zap.String("path", path),
		zap.String("service", string(serviceName)),
		zap.String("rpc", string(methodName)),
	)
}
