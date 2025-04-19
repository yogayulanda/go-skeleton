// File: main.go

package main

import (
	"log"

	cmd "github.com/yogayulanda/go-skeleton/cmd"
	logging "github.com/yogayulanda/go-skeleton/pkg/logger"
)

func main() {

	defer logging.SyncLogger()
	// Inisialisasi dan eksekusi root command
	if err := cmd.Execute(); err != nil {
		log.Fatal("Error executing command:", err)
	}

}
