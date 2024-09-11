package main

import (
	"log"

	"github.com/Amman30/EnvPilot/cmd/cli"
	"github.com/Amman30/EnvPilot/pkg/pilot"
)

func main() {
	pilot.SetEnv(".env")
	rootCmd := cli.NewCmdRoot() 
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
