package main

import (
	"log"

	"github.com/Amman30/EnvPilot/cmd/cli"
)

func main() {
	
	rootCmd := cli.NewCmdRoot()
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}


}
