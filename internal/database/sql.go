package database

// Package database menyediakan fungsi untuk menginisialisasi koneksi ke database
// dan mengelola pengaturan koneksi menggunakan GORM.
// Package database provides functions to initialize a connection to the database
// and manage connection settings using GORM.
// Package database provides functions to initialize a connection to the database
// and manage connection settings using GORM.

// import (
// 	"fmt"
// 	"time"

// 	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/config"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // NewMySQLConnection menginisialisasi koneksi MySQL menggunakan GORM
// func NewMySQLConnection(cfg config.DB) (*gorm.DB, error) {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
// 		cfg.SqlServer.Username, cfg.SqlServer.Password, cfg.SqlServer.Host, cfg.SqlServer.Port, cfg.SqlServer.Name,
// 	)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return nil, err
// 	}
// 	sqlDB.SetMaxOpenConns(int(cfg.SqlServer.Maxconns))
// 	sqlDB.SetMaxIdleConns(int(cfg.SqlServer.Maxidleconns))
// 	sqlDB.SetConnMaxLifetime(time.Hour)

// 	return db, nil
// }
