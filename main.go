// File: main.go

package main

import (
	cmd "github.com/yogayulanda/go-skeleton/cmd"
)

func main() {
	// Execute the root command
	// This will start the server and handle any subcommands
	// such as "config" or "server"
	cmd.Execute()

}
