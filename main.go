package main

import (
	"log"

	"github.com/Amman30/EnvPilot/config"
)

func main() {
	
	config.Init(".env.example")

	port, err := config.Env.GetAsInt("PORT", 8080)
	if err != nil {
		log.Fatalf("Error retrieving PORT: %v", err)
	}
	log.Printf("Port value: %d", port)


	debug, err := config.Env.GetAsBool("DEBUG", false)
	if err != nil {
		log.Printf("Error retrieving DEBUG: %v", err)
	}
	log.Printf("Debug value: %t", debug)

	// Example 3: Retrieve a string value with a default
	dbURL, err := config.Env.GetAsString("DATABASE_URL", "localhost")
	if err != nil {
		log.Printf("Error retrieving DATABASE_URL: %v", err)
	}
	log.Printf("Database URL: %s", dbURL)

	// Example 4: Retrieve a float value with a default
	threshold, err := config.Env.GetAsFloat("THRESHOLD", 0.1)
	if err != nil {
		log.Printf("Error retrieving THRESHOLD: %v", err)
	}
	log.Printf("Threshold value: %f", threshold)

	// Example 5: Retrieve any type with a specified type
	dbURLAny, err := config.Env.GetAsAny("DATABASE_URL", "string", "localhost")
	if err != nil {
		log.Printf("Error retrieving DATABASE_URL: %v", err)
	}
	log.Printf("Database URL (any type): %s", dbURLAny.(string))

	// Example 6: Retrieve an integer without providing a default value (expect error if not found)
	portNoDefault, err := config.Env.GetAsInt("PORT_NO_DEFAULT")
	if err != nil {
		log.Printf("Error retrieving PORT_NO_DEFAULT: %v", err)
	}
	log.Printf("Port with no default: %d", portNoDefault)
}
