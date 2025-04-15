package main

import (
	"os"

	"github.com/yogayulanda/go-skeleton/internal/config"
	"github.com/yogayulanda/go-skeleton/internal/logging"

	"github.com/yogayulanda/go-skeleton/cmd/server"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	rootCMD = &cobra.Command{
		Use:   "go-skeleton",
		Short: "go-skeleton service",
		Long:  "A command-line interface for managing and running the go-skeleton service.",
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
	// Initialize configuration using Viper
	cobra.OnInitialize(config.Init)

	// Register subcommands
	rootCMD.AddCommand(configCMD)
	rootCMD.AddCommand(serverCMD)

	// Execute the root command
	if err := rootCMD.Execute(); err != nil {
		logging.Log.Fatal("Command execution failed", zap.Error(err))
		os.Exit(1)
	}
}
