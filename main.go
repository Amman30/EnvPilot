package main

import (
	"env-manager/config"

	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize the configuration by loading environment variables from a specified file
	config.Init(".env.example")

	// Create and configure a logger
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	// Example 1: Retrieve an integer value with a default
	port, err := config.Env.GetAsInt("PORT", 8080)
	if err != nil {
		logger.Fatalf("Error retrieving PORT: %v", err)
	}
	logger.Infof("Port value: %d", port)

	// Example 2: Retrieve a boolean value with a default
	debug, err := config.Env.GetAsBool("DEBUG", false)
	if err != nil {
		logger.Fatalf("Error retrieving DEBUG: %v", err)
	}
	logger.Infof("Debug value: %t", debug)

	// Example 3: Retrieve a string value with a default
	dbURL, err := config.Env.GetAsString("DATABASE_URL", "localhost")
	if err != nil {
		logger.Fatalf("Error retrieving DATABASE_URL: %v", err)
	}
	logger.Infof("Database URL: %s", dbURL)

	// Example 4: Retrieve a float value with a default
	threshold, err := config.Env.GetAsFloat("THRESHOLD", 0.1)
	if err != nil {
		logger.Fatalf("Error retrieving THRESHOLD: %v", err)
	}
	logger.Infof("Threshold value: %f", threshold)

	// Example 5: Retrieve any type with a specified type
	dbURLAny, err := config.Env.GetAsAny("DATABASE_URL", "string", "localhost")
	if err != nil {
		logger.Fatalf("Error retrieving DATABASE_URL: %v", err)
	}
	logger.Infof("Database URL (any type): %s", dbURLAny.(string))

	// Example 6: Retrieve an integer without providing a default value (expect error if not found)
	portNoDefault, err := config.Env.GetAsInt("PORT_NO_DEFAULT")
	if err != nil {
		logger.Fatalf("Error retrieving PORT_NO_DEFAULT: %v", err)
	}
	logger.Infof("Port with no default: %d", portNoDefault)
}
