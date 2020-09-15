package snek

import (
	"os"
	"testing"
)

func TestSnekDefault(t *testing.T) {
	envVar := "SOME_VAR"
	value := "testvalue"
	// call Default()
	Default(envVar, value)

	// Get value, make sure default is returned
	v := Get(envVar)
	if v != value {
		t.Error("Default() failed to store and retreive default value")
	}
}

func TestSnekGet(t *testing.T) {
	envVar := "ANOTHER_VAR"
	value := "testvalue"
	// Get nonexistent value, make sure it returns ""
	v := Get(envVar)
	if v != "" {
		t.Error("Get() failed to return empty string when value not present")
	}
	// Set default
	Default(envVar, value)

	// Get value again, make sure default is returned
	v1 := Get(envVar)
	if v1 != value {
		t.Error("Get() failed to return default value when set")
	}

	// Set ENV_VAR
	os.Setenv(envVar, "newvalue")

	// get value again, make sure env var value is returned'
	v2 := Get(envVar)
	if v2 != "newvalue" {
		t.Error("Get() failed to override default value with env var when set")
	}
}

func TestSnekGetOrError(t *testing.T) {
	envVar := "YET_ANOTHER_VAR"
	value := "testvalue"
	// Get nonexistent value, make sure error is returned
	v, err := GetOrError(envVar)
	if v != "" {
		t.Error("GetOrError() failed to return empty string when value not present")
	}
	if err == nil {
		t.Error("GetOrError() failed to return error when value not present")
	}

	// Set default
	Default(envVar, value)

	// Get value again, make sure default is returned
	v1, err := GetOrError(envVar)
	if v1 != value {
		t.Error("GetOrError() failed to return default value when set")
	}

	if err != nil {
		t.Error("GetOrError() returned an error when default value was set and returned successfully")
	}

	// Set ENV_VAR
	os.Setenv(envVar, "newvalue")

	// get value again, make sure env var value is returned
	v2, err := GetOrError(envVar)
	if v2 != "newvalue" {
		t.Error("GetOrError() failed to override default value with env var when set")
	}

	if err != nil {
		t.Error("GetOrError() returned an error when env var value was set and returned successfully")
	}

}
