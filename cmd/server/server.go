package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/config"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/di"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/protocol/grpc"
	grpcgateway "gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/protocol/grpc-gateway"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/utils"
)

func RunServer() {
	// Load config
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	// Init DI container
	container := di.InitContainer(cfg)

	// Run gRPC Server
	go func() {
		grpcServer := grpc.NewGRPCServer(container)
		if err := grpc.StartGRPCServer(container.Config.GRPCPORT, grpcServer); err != nil {
			panic(err)
		}
	}()

	// Run gRPC-Gateway Server (REST Proxy)
	go func() {
		if err := grpcgateway.RunServerGrpcGW(context.Background(), container); err != nil {
			panic(err)
		}
	}()

	utils.LogAvailableEndpoints(container.Logger)
	// Graceful shutdown handling
	waitForShutdown()
}

// waitForShutdown handles OS signals and performs graceful shutdown if needed
func waitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
}
