package pilot

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fsnotify/fsnotify"
)



type EnvStore struct {
	variables map[string]interface{}
	filePath  string
	watcher   *fsnotify.Watcher
}

// Global instance to hold environment variables
var Env *EnvStore

// Init initializes the EnvStore by loading environment variables from the specified file.
// It also starts a watcher to reload the file when it changes.
func SetEnv(filename string) {
	if filename == "" {
		filename = ".env"
	}

	// Initialize EnvStore
	Env = &EnvStore{
		variables: make(map[string]interface{}),
		filePath:  filename,
	}
	Env.loadEnvFile()

	// Start file watcher
	if err := Env.startWatching(); err != nil {
		log.Fatalf("Error starting file watcher: %s", err)
	}
}

// loadEnvFile reads environment variables from a specified file and populates the EnvStore.
func (e *EnvStore) loadEnvFile() {
	file, err := os.Open(e.filePath)
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
		e.variables[key] = value
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading .env file: %s", err)
	}
}

// startWatching sets up a file watcher to reload the environment variables when the file changes.
func (e *EnvStore) startWatching() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	e.watcher = watcher

	err = watcher.Add(e.filePath)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Environment file changed. Reloading...")
					e.loadEnvFile()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("Error watching file: %s", err)
			}
		}
	}()

	return nil
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
// the variable is not found or cannot be converted to the requested
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


func (e *EnvStore) SetEnvValue(key string, value string, valueType string, filename string) error {
	switch valueType {
	case "string":
		e.variables[key] = value
	case "int":
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("invalid int value: %s", err)
		}
		e.variables[key] = intValue
	case "bool":
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("invalid bool value: %s", err)
		}
		e.variables[key] = boolValue
	case "float":
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("invalid float value: %s", err)
		}
		e.variables[key] = floatValue
	default:
		return fmt.Errorf("unsupported value type: %s", valueType)
	}

	// Save the updated environment variable to the file
	return e.saveToFile(key, value, filename)
}

// saveToFile updates the environment file with the new key-value pair
func (e *EnvStore) saveToFile(key, value string, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening file %s: %s", filename, err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s=%s\n", key, value))
	if err != nil {
		return fmt.Errorf("error writing to file %s: %s", filename, err)
	}
	return nil
}