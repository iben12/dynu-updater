package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
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

func TestGetIP(t *testing.T) {
	// Create a test server to simulate the API response
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := IPResponse{
			Ip: "127.0.0.1",
		}

		json.NewEncoder(w).Encode(response)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	ipAPIURL := server.URL

	ip := getIP(ipAPIURL)

	if ip != "127.0.0.1" {
		t.Errorf("Expected IP: 127.0.0.1, got: %s", ip)
	}
}

func TestUpdateIP(t *testing.T) {
	// Create a test server to simulate the API response
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	config := UpdateConfig{
		User:      "testuser",
		Password:  "testpass",
		Domain:    "example.com",
		ServerURL: server.URL,
	}

	err := updateIp(config, "127.0.0.1")

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestDoUpdate(t *testing.T) {
	// Create a test server to simulate the API responses
	ipHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := IPResponse{
			Ip: "127.0.0.1",
		}

		json.NewEncoder(w).Encode(response)
	})

	dynuHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("nochg"))
	})

	ipServer := httptest.NewServer(ipHandler)
	defer ipServer.Close()

	dynuServer := httptest.NewServer(dynuHandler)
	defer dynuServer.Close()

	config := UpdateConfig{
		User:      "testuser",
		Password:  "testpass",
		Domain:    "example.com",
		ServerURL: dynuServer.URL,
		IpServer:  ipServer.URL,
	}

	// Create a buffer to capture the program's output
	var outputBuffer bytes.Buffer

	oldLogger := logger
	logger = log.New(&outputBuffer, "", 5)

	// Call the doUpdate function
	doUpdate(config)

	// Check if the expected IP and update message are present in the output
	expectedIP := "Current IP is 127.0.0.1"
	if !bytes.Contains(outputBuffer.Bytes(), []byte(expectedIP)) {
		t.Errorf("Expected output to contain: %s", expectedIP)
	}

	expectedUpdateMsg := "Update result is: nochg"
	if !bytes.Contains(outputBuffer.Bytes(), []byte(expectedUpdateMsg)) {
		t.Errorf("Expected output to contain: %s", expectedUpdateMsg)
	}

	logger = oldLogger
}
