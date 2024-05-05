package main

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// Set up the environment variables for testing
	os.Setenv("API_KEY", "secret")
	os.Setenv("DNS_ID", "11111")
	os.Setenv("DOMAIN", "example.com")
	os.Setenv("PERIOD_HOURS", "1")

	expectedConfig := Config{
		ApiKey: "secret",
		DnsId:  "11111",
		Domain: "example.com",
		Period: 1,
	}

	config := getConfig()

	if config != expectedConfig {
		t.Errorf("Expected: %s, got: %s", expectedConfig.ApiKey, config.ApiKey)
	}

	// Clean up the environment variables after the test
	os.Unsetenv("API_KEY")
	os.Unsetenv("DNS_ID")
	os.Unsetenv("DOMAIN")
	os.Unsetenv("PERIOD_HOURS")
}
