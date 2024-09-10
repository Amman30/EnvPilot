# env-manager

**`env-manager`** is a Go package designed for managing environment variables in your Go applications. It offers a simple and flexible way to load and access environment variables with support for default values and type safety.

**Installation**

To install **`env-manager`**, use the following `go get` command:

```sh
go get github.com/Amman30/env-manager


Code Snippet ```bash
package main

import (
	"log"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/env-manager/config"
)

func main() {
	// Initialize the configuration by specifying the .env file
	config.Init(".env.example")

	// Create and configure a logger
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
}
```

# Retrieving Environment Variables

# GetAsString

```bash 
value, err := config.Env.GetAsString("KEY", "default_value")
if err != nil {
	log.Fatalf("Error retrieving KEY: %v", err)
}
log.Infof("Value: %s", value)
```


# GetAsInt
```bash
value, err := config.Env.GetAsInt("KEY", 8080)
if err != nil {
	log.Fatalf("Error retrieving KEY: %v", err)
}
log.Infof("Value: %d", value)

```

# GetAsBool

```bash 
value, err := config.Env.GetAsBool("KEY", false)
if err != nil {
	log.Fatalf("Error retrieving KEY: %
```
