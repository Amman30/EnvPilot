# EnvPilot

**`EnvPilot`**  is a Go package and CLI tool for managing environment variables. With EnvPilot, you can easily set, retrieve, and manage environment variables both programmatically and through the command line.It offers a simple and flexible way to load and access environment variables with support for default values and type safety.


**Installation**

To install **`EnvPilot`**, use the following `go get` command:

```sh
go get github.com/Amman30/EnvPilot@v0.1.3
```

# CLI Usage

Once EnvPilot is installed, you can use the CLI to manage environment variables. Here are some examples:

# Command
```bash
go install github.com/Amman30/EnvPilot/cmd/cli@0.1.1
```

# Set Environment Variables

To set an environment variable, use the set command. You can specify the variable type and the file to save it to.

Example to set an integer variable:
``` bash
pilot set MY_VAR=123222 --type int --file .env
```

Example to set an string variable in .env.example:
``` bash
pilot set GREETING=Hello --type string --file .env.example
```

# Retrieving Environment Variables
```bash
package main

import (
    "fmt"
    "log"

    "github.com/Amman30/EnvPilot/pkg/pilot"
)

func main() {
    // Initialize environment from .env file
    err := pilot.SetEnv(".env")
    if err != nil {
        log.Fatalf("Error initializing environment: %v", err)
    }
}
```


## GetAsString

The `GetAsString` method retrieves a string value from the environment variables. You can also provide a default value to return if the variable is not found.

### Arguments
- `key` (string): The environment variable key to retrieve.
- `defaultValue` (optional string): A default string value to return if the environment variable is not found. If not provided, an error will be returned if the key is not found.

### Example: Retrieve a String

```go
value, err := pilot.Env.GetAsString("KEY", "default_value")
if err != nil {
	log.Fatalf("Error retrieving KEY: %v", err)
}
log.Infof("Value: %s", value)


# GetAsInt
```bash
value, err := pilot.Env.GetAsInt("KEY", 8080)
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
value, err := pilot.Env.GetAsBool("KEY", false)
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
value, err := pilot.Env.GetAsFloat("KEY", 0.1)
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
value, err := pilot.Env.GetAsAny("KEY", "int", 8080)
if err != nil {
	log.Fatalf("Error retrieving KEY: %v", err)
}
log.Infof("Value: %d", value.(int))

```


# Dynamic Configuration Reloading

`EnvPilot` includes support for dynamic configuration reloading using the `fsnotify` package. This allows your application to automatically reload environment, making it easier to adapt to configuration updates without restarting the application.

To enable dynamic reloading, simply ensure that `fsnotify` is included in your project dependencies. The package will monitor the specified file path for changes and reload the configuration as needed.

## License

This project is licensed under the [MIT License](LICENSE).
