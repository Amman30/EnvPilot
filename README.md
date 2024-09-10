# EnvPilot

**`EnvPilot`** is a Go package designed for managing environment variables in your Go applications. It offers a simple and flexible way to load and access environment variables with support for default values and type safety.

**Installation**

To install **`EnvPilot`**, use the following `go get` command:

```sh
go get github.com/Amman30/EnvPilot
```

# Retrieving Environment Variables
```bash
package main

import (
	"github.com/Amman30/EnvPilot/config"
)

func main() {
	// Initialize the configuration by specifying the .env file
	config.Init(".env.example")   // you can also use .env.production , if not provided .env is default value
}
```


## GetAsString

The `GetAsString` method retrieves a string value from the environment variables. You can also provide a default value to return if the variable is not found.

### Arguments
- `key` (string): The environment variable key to retrieve.
- `defaultValue` (optional string): A default string value to return if the environment variable is not found. If not provided, an error will be returned if the key is not found.

### Example: Retrieve a String

```go
value, err := config.Env.GetAsString("KEY", "default_value")
if err != nil {
	log.Fatalf("Error retrieving KEY: %v", err)
}
log.Infof("Value: %s", value)


# GetAsInt
```bash
value, err := config.Env.GetAsInt("KEY", 8080)
if err != nil {
	log.Fatalf("Error retrieving KEY: %v", err)
}
log.Infof("Value: %d", value)

```

## GetAsBool

The `GetAsBool` method retrieves a boolean value from the environment variables. You can also provide a default value to return if the variable is not found.

### Arguments
- `key` (string): The environment variable key to retrieve.
- `defaultValue` (optional bool): A default boolean value to return if the environment variable is not found. If not provided, an error will be returned if the key is not found.

### Example: Retrieve a Boolean

```go
value, err := config.Env.GetAsBool("KEY", false)
if err != nil {
	log.Fatalf("Error retrieving KEY: %v", err)
}
log.Infof("Value: %t", value)
```

## GetAsFloat

The `GetAsFloat` method retrieves a float value from the environment variables. You can also provide a default value to return if the variable is not found.

### Arguments
- `key` (string): The environment variable key to retrieve.
- `defaultValue` (optional float64): A default float value to return if the environment variable is not found. If not provided, an error will be returned if the key is not found.

### Example: Retrieve a Float

```go
value, err := config.Env.GetAsFloat("KEY", 0.1)
if err != nil {
	log.Fatalf("Error retrieving KEY: %v", err)
}
log.Infof("Value: %f", value)
```
# GetAsAny

The `GetAsAny` method retrieves a value of any type from the environment variables based on the specified target type. It also allows you to provide a default value if the variable is not found or cannot be converted to the requested type.

### Arguments
- `key` (string): The environment variable key to retrieve.
- `targetType` (string): The type to which the value should be converted. Supported types include `"string"`, `"int"`, `"bool"`, and `"float"`.
- `defaultValue` (optional): A default value to return if the environment variable is not found or cannot be converted.

### Example: Retrieve an Integer

```go
value, err := config.Env.GetAsAny("KEY", "int", 8080)
if err != nil {
	log.Fatalf("Error retrieving KEY: %v", err)
}
log.Infof("Value: %d", value.(int))

