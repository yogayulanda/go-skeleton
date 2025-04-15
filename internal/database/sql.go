package database

import (
	"fmt"

	"github.com/yogayulanda/go-skeleton/internal/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewSQLServerConnection membuat koneksi ke database SQL Server
func NewSQLServerConnection(cfg *config.App) (*gorm.DB, error) {
	// Format string DSN untuk SQL Server
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		cfg.MSSQLUSER, cfg.MSSQLPASSWORD, cfg.MSSQLHOST, cfg.MSSQLPORT, cfg.MSSQLDB)

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
