package snek

import (
	"fmt"
	"os"
)

/* Snek - a very simple environment variable-based config management system, focused on
container/kubernetes use cases. Very small fangs. */

var defaults = make(map[string]string)

// Default registers a default value for a particular config environment variable
func Default(envvar string, defaultValue string) {
	defaults[envvar] = defaultValue
}

// Get retrieves the value for a config key. If the env var is set, it is retrieved.
// Otherwise, the default value is returned. If neither exist, an empty string "" is returned.
func Get(envvar string) string {
	// return env var if it exists
	envval := os.Getenv(envvar)
	if envval != "" {
		return envval
	}

	// otherwise return the default or "" if no default was registered
	val, _ := defaults[envvar]
	return val
}

// GetOrError retrieves the value for a config key. If the env var is set, it is retrieved.
// Otherwise, the default value is returned. If neither exist, an error is thrown.
func GetOrError(envvar string) (string, error) {
	// return env var if it exists
	envVal := os.Getenv(envvar)
	if envVal != "" {
		return envVal, nil
	}

	// otherwise return the default or "" if no default was registered
	defaultVal, present := defaults[envvar]
	if present {
		return defaultVal, nil
	}

	return "", fmt.Errorf("Config key %s not set, either as environment variable or default value", envvar)
}
