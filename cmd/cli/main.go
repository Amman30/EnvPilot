package main

import (
	"log"
)

func main() {
	
	rootCmd := NewCmdRoot()
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
