package main

import (
	"os"

	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/config"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/logging"

	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/cmd/server"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	rootCMD = &cobra.Command{
		Use:   "if-trx-history",
		Short: "if-trx-history service",
		Long:  "A command-line interface for managing and running the if-trx-history service.",
	}

	configCMD = &cobra.Command{
		Use:   "config",
		Short: "Show configuration settings",
		Run: func(*cobra.Command, []string) {
			config.Show()
		},
	}

	serverCMD = &cobra.Command{
		Use:   "server",
		Short: "Run the gRPC and HTTP servers",
		Run: func(*cobra.Command, []string) {
			server.RunServer()
		},
	}
)

func main() {
	// Initialize logging

	logging.InitLogger(false)
	defer logging.SyncLogger()

	// Initialize configuration using Viper
	cobra.OnInitialize(config.Init)

	// Register subcommands
	rootCMD.AddCommand(configCMD)
	rootCMD.AddCommand(serverCMD)

	// Execute the root command
	if err := rootCMD.Execute(); err != nil {
		logging.Logger.Fatal("Command execution failed", zap.Error(err))
		os.Exit(1)
	}
}
