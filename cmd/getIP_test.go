package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
