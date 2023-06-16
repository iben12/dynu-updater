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
	os.Setenv("PASSWORD_HASH", "testhash")
	os.Setenv("DOMAIN", "example.com")
	os.Setenv("PERIOD_HOURS", "1")

	expectedConfig := Config{
		User:   "testuser",
		Secret: "testhash",
		Domain: "example.com",
		Period: 1,
	}

	config := getConfig()

	if config != expectedConfig {
		t.Errorf("Expected secret: %s, got: %s", expectedConfig.Secret, config.Secret)
	}

	os.Unsetenv("PASSWORD_HASH")

	config = getConfig()

	expectedConfig.Secret = fmt.Sprintf("%v", sha256.Sum256([]byte("testpass")))

	if config != expectedConfig {
		t.Errorf("Expected secret: %s, got: %s", expectedConfig.Secret, config.Secret)
	}

	// Clean up the environment variables after the test
	os.Unsetenv("USERNAME")
	os.Unsetenv("PASSWORD")
	os.Unsetenv("DOMAIN")
	os.Unsetenv("PERIOD_HOURS")
}
