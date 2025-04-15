// File: main.go

package main

import (
	"github.com/yogayulanda/go-skeleton/cmd/server"
)

func main() {
	// Execute the root command
	// This will start the server and handle any subcommands
	// such as "config" or "server"
	server.Execute()
}
