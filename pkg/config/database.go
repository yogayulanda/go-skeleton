package config

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDB(cfg *App) (*gorm.DB, error) {
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
	return db, nil
}
