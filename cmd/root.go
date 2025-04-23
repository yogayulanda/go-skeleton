package server

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yogayulanda/go-skeleton/pkg/config"
	"github.com/yogayulanda/go-skeleton/pkg/di"
	logging "github.com/yogayulanda/go-skeleton/pkg/logger"
	"github.com/yogayulanda/go-skeleton/pkg/utils"
	"go.uber.org/zap"
)

// Root command
var rootCMD = &cobra.Command{
	Use:   "go-skeleton",
	Short: "go-skeleton service",
	Long:  "A command-line interface for managing and running the go-skeleton service.",
}

// Config command
var configCMD = &cobra.Command{
	Use:   "config",
	Short: "Show configuration settings",
	Run: func(cmd *cobra.Command, args []string) {
		config.ShowConfig()
	},
}

// Rollback command for rolling back migrations
var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback the last applied migration",
	Long:  `This command will rollback the last applied migration using Goose.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the DI container and config
		// Memuat konfigurasi
		cfg, err := config.LoadConfig()
		if err != nil {
			panic(fmt.Errorf("error loading config: %v", err))
		}

		logging.InitLogger(cfg.LOG_LEVEL)
		log := logging.Log

		db, err := config.InitDB(cfg, log)
		if err != nil {
			log.Fatal("failed to init DB", zap.Error(err))
		}
		// Perform rollback
		err = config.RollbackSQLMigrations(db) // Call to rollback migrations
		if err != nil {
			log.Fatal("Failed to rollback migrations", zap.Error(err))
		}

		fmt.Println("Migrations rolled back successfully")
	},
}

// Server command
var serverCMD = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC and HTTP servers",
	Run: func(cmd *cobra.Command, args []string) {
		// Memuat konfigurasi
		cfg, err := config.LoadConfig()
		if err != nil {
			panic(fmt.Errorf("error loading config: %v", err))
		}
		// Menyiapkan DI container
		container := di.InitContainer(cfg)
		container.Log.Info("🚀 Starting the server...")
		// Menjalankan server dan menangani error jika ada
		if err := RunServer(container, cfg); err != nil {
			log.Fatalf("Error running server: %v", err)
		}
	},
}

// Execute menjalankan root command dan perintah-perintah lainnya
func Execute() error {

	// Menyiapkan konfigurasi dan container DI
	if err := config.Init(); err != nil {
		panic(fmt.Errorf("failed to initialize configuration: %v", err))
	}

	// Remove soon
	token, err := utils.GenerateToken("user123", "admin")
	if err != nil {
		log.Fatalf("Error generating token: %v", err)
	}
	fmt.Println("Generated Token:", token)

	// Menambahkan command
	rootCMD.AddCommand(configCMD) // Menambahkan perintah config
	rootCMD.AddCommand(serverCMD) // Menambahkan perintah server
	rootCMD.AddCommand(rollbackCmd)

	// Eksekusi command yang dipilih
	if err := rootCMD.Execute(); err != nil {
		panic(fmt.Errorf("error executing root command: %v", err))
	}
	return nil
}
