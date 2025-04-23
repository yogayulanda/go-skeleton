package config

import (
	"fmt"

	"github.com/pressly/goose"
	"github.com/yogayulanda/go-skeleton/pkg/models"
	"go.uber.org/zap"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDB(cfg *App, log *zap.Logger) (*gorm.DB, error) {
	// Format DSN (data source name) untuk SQL Server
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		cfg.MSSQL_USER,
		cfg.MSSQL_PASS,
		cfg.MSSQL_HOST,
		cfg.MSSQL_PORT,
		cfg.MSSQL_DB,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQL Server: %v", err)
	}

	// Perform auto-migration for the ErrorCode model (or other models)
	if err := db.AutoMigrate(&models.ErrorCode{}); err != nil {
		return nil, fmt.Errorf("auto-migration failed: %v", err)
	}

	// Run Goose migration (using SQL Server driver and migration folder)
	if err := runMigrations(db, cfg.MIGRATIONS_FOLDER); err != nil {
		return nil, fmt.Errorf("migration failed: %v", err)
	}

	log.Info("Migrations completed successfully")
	return db, nil
}

// runMigrations will run the Goose migrations on the given database connection
func runMigrations(db *gorm.DB, migrationsFolder string) error {

	// Get raw *sql.DB from GORM instance
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get *sql.DB from GORM DB: %v", err)
	}

	// Open the database connection with Goose (it will use *sql.DB)
	if err := goose.SetDialect("mssql"); err != nil {
		return fmt.Errorf("failed to set Goose dialect: %v", err)
	}

	// Run the migrations
	if err := goose.Up(sqlDB, migrationsFolder); err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	return nil
}

// rollbackSQLMigrations will execute SQL migrations rollback (undo the applied changes)
func RollbackSQLMigrations(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get *sql.DB from GORM DB: %v", err)
	}
	// Set Goose dialect for SQL Server
	if err := goose.SetDialect("mssql"); err != nil {
		return fmt.Errorf("failed to set Goose dialect: %v", err)
	}

	// Rollback the last migration
	migrationsFolder := "./migrations"
	if err := goose.Down(sqlDB, migrationsFolder); err != nil {
		return fmt.Errorf("failed to rollback migration: %v", err)
	}

	return nil
}
