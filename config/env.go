package config

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type EnvStore struct {
	variables map[string]interface{}
}

// Global instance to hold environment variables
var Env = &EnvStore{variables: make(map[string]interface{})}

// init initializes the EnvStore by loading environment variables from the default ".env" file.
func Init(filename string) {
    if filename == "" {
        filename = ".env"
    }
    loadEnvFile(filename)
}

// loadEnvFile reads environment variables from a specified file and populates the EnvStore.
// It opens the file, reads it line by line, and parses key-value pairs to store in the EnvStore.
// Lines starting with "#" or that are empty are ignored. The function logs and terminates the program
// if the file cannot be opened or read.
func loadEnvFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening .env file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore comments and empty lines
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		// Split the line into key-value pairs
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Store the key-value pair in the Variables map
		Env.variables[key] = value
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading .env file: %s", err)
	}
}

// GetAsString retrieves the value of an environment variable as a string. If the variable is not found
// and a default value is provided, the default value is returned. If no default value is provided and
// the variable is not found, an error is returned.
func (e *EnvStore) GetAsString(key string, defaultValue ...string) (string, error) {
	if value, exists := e.variables[key]; exists {
		if str, ok := value.(string); ok {
			return str, nil
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0], nil
	}
	return "", errors.New("variable not found or type mismatch and no default value provided")
}

// GetAsInt retrieves the value of an environment variable as an integer. If the variable is not found
// and a default value is provided, the default value is returned. If no default value is provided and
// the variable is not found or cannot be converted to an integer, an error is returned.
func (e *EnvStore) GetAsInt(key string, defaultValue ...int) (int, error) {
	if value, exists := e.variables[key]; exists {
		switch v := value.(type) {
		case string:
			i, err := strconv.Atoi(v)
			if err == nil {
				e.variables[key] = i // Update to store as int for future lookups
				return i, nil
			}
		case int:
			return v, nil
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0], nil
	}
	return 0, errors.New("variable not found or type mismatch and no default value provided")
}

// GetAsBool retrieves the value of an environment variable as a boolean. If the variable is not found
// and a default value is provided, the default value is returned. If no default value is provided and
// the variable is not found or cannot be converted to a boolean, an error is returned.
func (e *EnvStore) GetAsBool(key string, defaultValue ...bool) (bool, error) {
	if value, exists := e.variables[key]; exists {
		switch v := value.(type) {
		case string:
			b, err := strconv.ParseBool(v)
			if err == nil {
				e.variables[key] = b 
				return b, nil
			}
		case bool:
			return v, nil
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0], nil
	}
	return false, errors.New("variable not found or type mismatch and no default value provided")
}


// GetAsFloat retrieves the value of an environment variable as a float. If the variable is not found
// and a default value is provided, the default value is returned. If no default value is provided and
// the variable is not found or cannot be converted to a float, an error is returned.
func (e *EnvStore) GetAsFloat(key string, defaultValue ...float64) (float64, error) {
	if value, exists := e.variables[key]; exists {
		switch v := value.(type) {
		case string:
			f, err := strconv.ParseFloat(v, 64)
			if err == nil {
				e.variables[key] = f
				return f, nil
			}
		case float64:
			return v, nil
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0], nil
	}
	return 0.0, errors.New("variable not found or type mismatch and no default value provided")
}

// GetAsAny retrieves the value of an environment variable as any type specified by targetType.
// The targetType must be one of "string", "int", "bool", or "float". If the variable is not found and
// a default value is provided, the default value is returned. If no default value is provided and
// the variable is not found or cannot be converted to the requested type, an error is returned.
func (e *EnvStore) GetAsAny(key string, targetType string, defaultValue ...interface{}) (interface{}, error) {
    if value, exists := e.variables[key]; exists {
        switch targetType {
        case "string":
            if str, ok := value.(string); ok {
                return str, nil
            }
        case "int":
            switch v := value.(type) {
            case string:
                i, err := strconv.Atoi(v)
                if err == nil {
                    e.variables[key] = i
                    return i, nil
                }
            case int:
                return v, nil
            }
        case "bool":
            switch v := value.(type) {
            case string:
                b, err := strconv.ParseBool(v)
                if err == nil {
                    e.variables[key] = b
                    return b, nil
                }
            case bool:
                return v, nil
            }
        case "float":
            switch v := value.(type) {
            case string:
                f, err := strconv.ParseFloat(v, 64)
                if err == nil {
                    e.variables[key] = f
                    return f, nil
                }
            case float64:
                return v, nil
            }
        }
    }
    if len(defaultValue) > 0 {
        return defaultValue[0], nil
    }
    return nil, errors.New("variable not found or cannot be converted to the requested type")
}





