package config

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewSQLServerConnection membuat koneksi ke database SQL Server
func NewSQLServerConnection(cfg *App) (*gorm.DB, error) {
	// Format string DSN untuk SQL Server
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		cfg.MSSQL_USER, cfg.MSSQL_PASS, cfg.MSSQL_HOST, cfg.MSSQL_PORT, cfg.MSSQL_DB)

	// Membuka koneksi ke database SQL Server menggunakan GORM
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Pastikan koneksi ke database berhasil
	if err := db.Exec("SELECT 1").Error; err != nil {
		return nil, err
	}

	return db, nil
}
