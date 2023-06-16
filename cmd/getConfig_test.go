package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// Set up the environment variables for testing
	os.Setenv("USERNAME", "testuser")
	os.Setenv("PASSWORD", "testpass")
	os.Setenv("DOMAIN", "example.com")
	os.Setenv("PERIOD_HOURS", "1")

	expectedConfig := UpdateConfig{
		User:     "testuser",
		Password: fmt.Sprintf("%v", sha256.Sum256([]byte("testpass"))),
		Domain:   "example.com",
		Period:   1,
	}

	config := getConfig()

	if config.User != expectedConfig.User {
		t.Errorf("Expected user: %s, got: %s", expectedConfig.User, config.User)
	}

	// Clean up the environment variables after the test
	os.Unsetenv("USERNAME")
	os.Unsetenv("PASSWORD")
	os.Unsetenv("DOMAIN")
	os.Unsetenv("PERIOD_HOURS")
}
